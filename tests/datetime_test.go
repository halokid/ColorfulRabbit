package tests

import (
  "fmt"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "testing"
)

func TestGetNowMin(t *testing.T)  {
  m := GetNowMin()
  fmt.Println(m)
}

func TestGetMinBefore(t *testing.T) {
  /**
  now := time.Now()
  then1 := now.Add(time.Duration(-2) * time.Minute)
  fmt.Println(then1)
  //fmt.Println(then1.Format("2006-01-02 15:01"))
  fmt.Println(then1.Format("2006-01-02 15:04"))

  then := time.Duration(-2) * time.Minute
  fmt.Println(then)
  */
  m := GetMinBefore(2)
  fmt.Println(m)
}










