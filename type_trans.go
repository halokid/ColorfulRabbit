package ColorfulRabbit

import (
  "github.com/bitly/go-simplejson"
  "strconv"
)

/**
类型转换函数
 */

 // 字符串转json
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

// int 转 string
func IntToStr(i int) string {
   s := strconv.Itoa(i)
   return s
}

// string 转  float64
func StrToFlo64(s string) float64 {
  f, err := strconv.ParseFloat(s, 64)
  CheckError(err)
  return f
}

// int 转 float64
func IntToFlo64(i int) float64 {
  is := IntToStr(i)
  isf := StrToFlo64(is)
  return isf
}


// string 转 int
func StrToInt(s string) int {
  i, err := strconv.Atoi(s)
  CheckError(err)
  return i
}



















