package tests

import (
  "fmt"
  "github.com/r00tjimmy/ColorfulRabbit"
  "testing"
)

func TestRandInt64(t *testing.T)  {
  randx := ColorfulRabbit.RandInt(0, 5)
  fmt.Println(randx)
}