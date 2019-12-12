package ColorfulRabbit

import (
  "strings"
  "time"
)

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

func IndexOfI(element int, data []int) (int) {
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

func HasKeyCom(key string, mapx map[string]interface{}) bool {
  if _, ok := mapx[key]; ok {
    return true
  } else {
    return false
  }
}



func HasKeyInt(key string, mapx map[string]map[string]int) bool {
  // 二维map是否包含key
  if _, ok := mapx[key]; ok {
    return true
  } else {
    return false
  }
}


func HasKeyIntSig(key string, mapx map[string]int) bool {
  // 一维map是否包含key
  if _, ok := mapx[key]; ok {
    return true
  } else {
    return false
  }
}


func Contain(checkStr string, patt string) bool {
  // 是否包含某字符串
  if strings.Index(checkStr, patt) == -1 {
    return false
  } else {
    return true
  }
}


func DurTime(t1 time.Time, t2 time.Time) time.Duration {
  tDur := t2.Sub(t1)
  return tDur
}













