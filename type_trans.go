package ColorfulRabbit

import (
  "fmt"
  "github.com/bitly/go-simplejson"
  "strconv"
)

/**
类型转换函数
 */

func StrToJson(s string) *simplejson.Json {
  // 字符串转json
  sJs, err := simplejson.NewJson([]byte(s))
  CheckFatal(err, "string change to json error")
  return sJs
}


func IntToInt64(i int) int64 {
  // int转int64
  iStr := strconv.Itoa(i)   // int to string
  iInt64, err := strconv.ParseInt(iStr, 10, 64)   // string to int64
  CheckError(err)
  return iInt64
}


func Flo64ToStr(f float64, prec int) string {
  // float64转string，保留几位小数
  fStr := strconv.FormatFloat(f, 'f', prec, 64)
  return fStr
}

func IntToStr(i int) string {
  // int 转 string
   s := strconv.Itoa(i)
   return s
}

func StrToFlo64(s string) float64 {
  // string 转  float64
  f, err := strconv.ParseFloat(s, 64)
  CheckError(err)
  return f
}

func IntToFlo64(i int) float64 {
  // int 转 float64
  is := IntToStr(i)
  isf := StrToFlo64(is)
  return isf
}


func StrToInt(s string) int {
  // string 转 int
  i, err := strconv.Atoi(s)
  CheckError(err)
  return i
}

func Flo64ToInt(f float64) int {
  // float64 转 int
  fStr := Flo64ToStr(f, 0)
  fInt, err := strconv.Atoi(fStr)
  if err != nil {
    fmt.Println("Flo64ToInt error")
    return 0
  }
  return fInt
}



















