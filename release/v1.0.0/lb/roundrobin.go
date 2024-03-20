package main

import (
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
)

/**
负载均衡算法--轮询调度
 */

// RR: 基于 权重round robin算法的接口
type RR interface {
  Next() interface{}
  Add(node interface{}, weight int)
  RemoveAll()
  Reset()
}

const (
  RR_NGINX = 0    // nginx算法
  RR_LVS = 1      // lvs算法
)

func NewRR(rtype int) RR {
  // 算法实现工厂类
  if rtype == RR_NGINX {
    return &WNGINX{}
  } else if rtype == RR_LVS {
    return &WLVS{}
  }
  return nil
}

// 节点结构
type WeightNginx struct {
  Node              interface{}
  Weight            int         // 节点的总体权重， 越大的话， 节点被重复选择的可能性越大, 这个值直接影响节点的重复概率
  CurrentWeight     int         // 决定判断这一次轮询是否选择该节点的依据， 要大于上一个best节点的CurrentWeight值才会被选中取代其，成为当前的best节点
  EffectiveWeight   int         // 节点的直接权重， 这个值越大，越容易被首先选中， 这个值直接影响节点的被选中概率
}

func (ww *WeightNginx) fail() {
  ww.EffectiveWeight -= ww.Weight
  if ww.EffectiveWeight < 0 {
    ww.EffectiveWeight = 0
  }
}

// nginx算法实现类
type WNGINX struct {
  nodes   []*WeightNginx      // 节点列表
  n       int                 // 节点的数量
}

func (w *WNGINX) Add(node interface{}, weight int) {
  // 增加权重节点
  weighted := &WeightNginx{
    Node:               node,
    Weight:             weight,
    EffectiveWeight:    weight,
    CurrentWeight:      0,        // 默认为0
  }
  w.nodes = append(w.nodes, weighted)
  w.n++
}

func (w *WNGINX) RemoveAll() {
  w.nodes = w.nodes[:0]     // 清空节点列表
  w.n = 0
}

func (w *WNGINX) Next() interface{} {
  // 下次轮询事件, 下次的轮询会根据算法取得某一个节点
  if w.n == 0 {
    return nil
  }
  if w.n == 1 {
    return w.nodes[0].Node
  }

  return nextWeightNode(w.nodes).Node
}

func nextWeightNode(nodes []*WeightNginx) (best *WeightNginx) {
  // todo: 下次轮询取得节点的算法
  total := 0

  for i := 0; i < len(nodes); i++ {
    w := nodes[i]

    if w == nil {
      continue
    }

    // 每当被选择一个 currentWeight就会增加一个 effectWeight 的值, currentWeight默认为0
    // todo: EffectiveWeight 就是一个表示更有可能调用到的参数， EffectiveWeight越大， 就越有可能被调用到, 假如EffectiveWeight小的某一次轮询的某个节点没有被调用到， 则EffectiveWeight加1， 那么下次这个节点就有更大可能被调用到， 所以会有A1位置的逻辑
    w.CurrentWeight += w.EffectiveWeight      // todo: A2
    total += w.EffectiveWeight

    // todo: Weight越大， 那么EffectiveWeight 就会越来越大， 则 A2 位置的 CurrentWeight 会更大， 则选择该节点的可能性更大
    if w.EffectiveWeight < w.Weight {     // todo: A1
      w.EffectiveWeight++
    }

    if best == nil || w.CurrentWeight > best.CurrentWeight {
      best = w
    }
  }

  if best == nil {
    return nil
  }
  // todo:  把当前节点的权重下降， 使得下次寻找节点的时候，  w.CurrentWeight > best.CurrentWeight不会命中， 不会选择到这次选中的节点， 达到轮询的效果
  best.CurrentWeight -= total
  return best
}

func (w *WNGINX) Reset() {
  for _, s := range w.nodes {
    s.EffectiveWeight = s.Weight
    s.CurrentWeight = 0
  }
}

// --------------------------------------------

// 节点结构
type WeightLvs struct {
  Node          interface{}
  Weight        int
}

type WLVS struct {
  // lvs算法实现
  nodes     []*WeightLvs
  n         int
  gcd       int     // 通用的权重因子
  maxW      int     // 最大权重
  i         int     // 被选择的次数
  cw        int     // 当前的权重值
}

func (w *WLVS) Next() interface{} {
  if w.n == 0 {
    return nil
  }

  if w.n == 1 {
    return w.nodes[0].Node
  }

  for {
    w.i = (w.i + 1) % w.n
    if w.i == 0 {
      w.cw = w.cw - w.gcd
      if w.cw <= 0 {
        w.cw = w.maxW
        if w.cw == 0 {
          return nil
        }
      }
    }

    if w.nodes[w.i].Weight >= w.cw {
      return w.nodes[w.i].Node
    }
  }
}

//增加权重节点
func (w *WLVS) Add(node interface{}, weight int) {
  weighted := &WeightLvs{Node: node, Weight: weight}
  if weight > 0 {
    if w.gcd == 0 {
      w.gcd = weight
      w.maxW = weight
      w.i = -1
      w.cw = 0
    } else {
      w.gcd = gcd(w.gcd, weight)
      if w.maxW < weight {
        w.maxW = weight
      }
    }
  }
  w.nodes = append(w.nodes, weighted)
  w.n++
}


func gcd(x, y int) int {
  var t int
  for {
    t = (x % y)
    if t > 0 {
      x = y
      y = t
    } else {
      return y
    }
  }
}

func (w *WLVS) RemoveAll() {
  w.nodes = w.nodes[:0]
  w.n = 0
  w.gcd = 0
  w.maxW = 0
  w.i = -1
  w.cw = 0
}

func (w *WLVS) Reset() {
  w.i = -1
  w.cw = 0
}


// --------------------
var rrx = NewRR(RR_NGINX)

type handle struct {
  addrs []string        // 节点列表
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // 实现handle自己的ServeHTTP方法， 承接 http.ListenAndServe 调用
  addr := rrx.Next().(string)
  remote, err := url.Parse("http://" + addr)
  if err != nil {
    panic(err)
  }

  proxy := httputil.NewSingleHostReverseProxy(remote)
  proxy.ServeHTTP(w, r)
}

func startServer() {
  // 被代理的服务器host和port
  h := &handle{}
  h.addrs = []string{"172.0.0.1:8080, 172.0.0.2:8080"}

  w := 1
  for _, e := range h.addrs {
    rrx.Add(e, w)
    w++
  }

  err := http.ListenAndServe(":28080", h)
  if err != nil {
    log.Fatalln("ListenAndServe fail: ", err)
  }
}

func main() {
  startServer()
}









