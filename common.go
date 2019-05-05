package ColorfulRabbit

import "strings"

/**
 常规的函数
 */


// 获取slice某元素的index
func IndexOf(element string, data []string) (int) {
  for k, v := range data {
    if element == v {
      return k
    }
  }
  return -1 //not found.
}


// 判断map中key是否存在
//func HasKey(key string, mapx map[string]interface{}) bool {
func HasKey(key string, mapx map[string]map[string]interface{}) bool {
  if _, ok := mapx[key]; ok {
    return true
  } else {
    return false
  }
}


// 是否包含某字符串
func Contain(checkStr string, patt string) bool {
  if strings.Index(checkStr, patt) == -1 {
    return false
  } else {
    return true
  }
}





