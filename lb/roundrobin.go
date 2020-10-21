package lb
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
  Weight            int
  CurrentWeight     int
  EffectiveWeight   int
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
  }
  w.nodes = append(w.nodes, weighted)
  w.n++
}

func (w *WNGINX) RemoveAll() {
  w.nodes = w.nodes[:0]     // 清空节点列表
  w.n = 0
}

func (w *WNGINX) Next() interface{} {
  // 下次轮询事件
  if w.n == 0 {
    return nil
  }
  if w.n == 1 {
    return w.nodes[0].Node
  }

  return nextWeightNode(w.nodes).Node
}

func nextWeightNode(nodes []*WeightNginx) (best *WeightNginx) {
  total := 0

  for i := 0; i < len(nodes); i++ {
    w := nodes[i]

    if w == nil {
      continue
    }

    w.CurrentWeight += w.EffectiveWeight
    total += w.EffectiveWeight
    if w.EffectiveWeight < w.Weight {
      w.EffectiveWeight++
    }

    if best == nil || w.CurrentWeight > best.CurrentWeight {
      best = w
    }
  }

  if best == nil {
    return nil
  }
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







