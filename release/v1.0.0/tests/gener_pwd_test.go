package tests

import (
  "fmt"
  "github.com/halokid/ColorfulRabbit"
  "math/rand"
  "testing"
  "time"
)

func TestGenerPwd(t *testing.T) {
 //随机种子
  rand.Seed(time.Now().UnixNano())
  //test1()
  passwd := ColorfulRabbit.GenerPwd("mix", 16)
  fmt.Println(passwd)


  p2 := ColorfulRabbit.GenerPwd("mix", 16)
  fmt.Println(p2)
  //fmt.Printf("length:%d charset:%s\n", length, charset)
}

func TestTmp2(t *testing.T) {
  var pwd []byte = make([]byte, 9)
  fmt.Println(pwd)
}