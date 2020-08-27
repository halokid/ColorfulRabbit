package ColorfulRabbit
/**
slice 操作相关函数
 */

func RmRepeatSlcStr(slc []string) []string {
  // slice去重, 时间换空间，双重循环来过滤重复元素
  res := make([]string, 0, len(slc))
  tmp := map[string]struct{}{}
  for _, item := range slc {
    if _, ok := tmp[item]; !ok {
      tmp[item] = struct{}{}      // 空 struct 不占内存空间
      res = append(res, item)
    }
  }
  return res
}

func RmRepeatSlcInter(slc []interface{}) []interface{} {
  // slice去重, 时间换空间，双重循环来过滤重复元素
  res := make([]interface{}, 0, len(slc))
  //tmp := map[interface{}]struct{}{}
  tmp := map[string]struct{}{}
  for _, item := range slc {
    if _, ok := tmp[item.(string)]; !ok {
      tmp[item.(string)] = struct{}{}      // 空 struct 不占内存空间
      res = append(res, item)
    }
  }
  return res
}


func Union(slice1, slice2 []string) []string {
  //求并集
  m := make(map[string]int)
  for _, v := range slice1 {
    m[v]++
  }

  for _, v := range slice2 {
    times, _ := m[v]
    if times == 0 {
      slice1 = append(slice1, v)
    }
  }
  return slice1
}

func Intersect(slice1, slice2 []string) []string {
  //求交集
  m := make(map[string]int)
  nn := make([]string, 0)
  for _, v := range slice1 {
    m[v]++
  }

  for _, v := range slice2 {
    times, _ := m[v]
    if times == 1 {
      nn = append(nn, v)
    }
  }
  return nn
}

func Difference(slice1, slice2 []string) []string {
  //求差集 slice1减去交集
  m := make(map[string]int)
  nn := make([]string, 0)
  inter := Intersect(slice1, slice2)
  for _, v := range inter {
    m[v]++
  }

  for _, value := range slice1 {
    times, _ := m[value]
    if times == 0 {
      nn = append(nn, value)
    }
  }
  return nn
}

func InSlice(find string, sl []string) bool {
  // 判断slice中是否存在
  for _, v := range sl {
    if find == v {
      return true
    }
  }
  return false
}



