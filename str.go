package ColorfulRabbit

import "strings"

/**
字符操作
 */

func GetSpIdx(s string, sep string, idx int) string {
  // 根据特征字符获取第几个split的值
  // s: tps=0&typ=py
  sp := strings.Split(s, sep)
  i := 0
  if idx < 0 {
    i = len(sp) + idx
    return sp[i]
  }
  return sp[idx]
}
