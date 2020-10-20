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
    return &RRNGINX
  } else if rtype == RR_LVS {
    return &RRLVS
  }
  return nil
}




