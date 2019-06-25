package tests

import (
  "fmt"
  "github.com/halokid/ColorfulRabbit"
  "testing"
  "time"
  "unsafe"
)

func TestDurTime(t *testing.T) {
  t1 := time.Now()
  time.Sleep(3 * time.Second)
  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}

func TestSizeof(t *testing.T) {
  i := 1
  fmt.Println(unsafe.Sizeof(&i))      // 8

  var s struct{}
  fmt.Println(unsafe.Sizeof(&s))      // 8
  fmt.Println(unsafe.Sizeof(s))       // 0

  //type a struct {}
  //fmt.Println(unsafe.Sizeof(&a))
}

func TestGenericFc(t *testing.T) {
  var i interface{}
  i = 5
  fmt.Printf("Type:  %T | Value: %v\n", i, i)

  switch o := i.(type) {
  case int:
    fmt.Printf("%5d\n", o)
  case float64:
    fmt.Printf("%7.3f\n", o)
  default:
    fmt.Printf("%+v\n", o)
  }
}



