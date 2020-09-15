package main

import "fmt"
import "time"

func main() {
  fmt.Println("run in main goroutine")
  n := 10
  for i:=0; i<n; i++ {
    go func() {
      fmt.Println("dead loop goroutine start")
    }()
  }
  for {
    time.Sleep(time.Second)
    fmt.Println("main goroutine running")
  }
}
