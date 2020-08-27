package ColorfulRabbit


/**
map数据类型使用的函数
 */

 type Xmap map[interface{}]interface{}

func (m *Xmap) Keys() []interface{} {
  // 获得map的key
  keys := make([]interface{}, 0, len(*m))
  for k := range *m {
    keys = append(keys, k)
  }
  return keys
}

func (m *Xmap) Items() []interface{} {
  // 获得map的val
  items := make([]interface{}, 0, len(*m))
  for _, v := range *m {
    items = append(items, v)
  }
  return items
}

 /**
func MakeMapIterKiVi(m map[int]int) {
  // map 的 key 和 val 都 interface{} 化
  // key type is int, val type is int
  im := make(map[interface{}]interface{})
  for i, v := range m {

  }
}
 */


func GetKeysSb(m map[string]bool) []string {
  // 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

func GetKeysSsi(m map[string]map[string]interface{}) []string {
  // 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}



















