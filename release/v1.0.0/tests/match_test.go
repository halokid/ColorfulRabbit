package tests

import (
  "fmt"
  "github.com/halokid/ColorfulRabbit"
  "testing"
)

func TestRandInt64(t *testing.T)  {
  randx := ColorfulRabbit.RandInt(0, 5)
  fmt.Println(randx)
}