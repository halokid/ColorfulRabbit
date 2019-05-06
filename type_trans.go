package ColorfulRabbit

import (
  "github.com/bitly/go-simplejson"
  "strconv"
)

/**
类型转换函数
 */

func StrToJson(s string) *simplejson.Json {
  sJs, err := simplejson.NewJson([]byte(s))
  CheckFatal(err, "string change to json error")
  return sJs
}


// int转int64
func IntToInt64(i int) int64 {
  iStr := strconv.Itoa(i)   // int to string
  iInt64, err := strconv.ParseInt(iStr, 10, 64)   // string to int64
  CheckError(err)
  return iInt64
}


// float64转string，保留几位小数
func Flo64ToStr(f float64, prec int) string {
  fStr := strconv.FormatFloat(f, 'f', prec, 64)
  return fStr
}
