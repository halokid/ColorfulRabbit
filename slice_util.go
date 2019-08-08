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


