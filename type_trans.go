package ColorfulRabbit

import "github.com/bitly/go-simplejson"

/**
类型转换函数
 */

func StrToJson(s string) *simplejson.Json {
  sJs, err := simplejson.NewJson([]byte(s))
  CheckFatal(err, "string change to json error")
  return sJs
}



