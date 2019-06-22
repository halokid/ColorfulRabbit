package tests

import (
  "fmt"
  "github.com/halokid/ColorfulRabbit"
  "testing"
  "time"
)

func TestDurTime(t *testing.T) {
  t1 := time.Now()
  time.Sleep(3 * time.Second)
  t2 := time.Now()
  t3 := ColorfulRabbit.DurTime(t1, t2)
  fmt.Println(t3)
}




