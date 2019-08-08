package ColorfulRabbit

import (
  "fmt"
  "testing"
)

func TestStrToJson(t *testing.T) {
  s :=  `["10.86.13.55:7084", "10.86.16.166:9001", "10.86.64.139:7082", "10.86.64.140:7082", "10.86.64.141:7083", "10.86.64.35:9550"]`
  sj := StrToJson(s)
  fmt.Println(sj)

  sjArr, _ := sj.Array()
  fmt.Println(sjArr)

  fmt.Println(sjArr[2])
}

