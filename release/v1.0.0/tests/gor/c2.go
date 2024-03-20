package main

import (
  "fmt"
  "runtime"
  "sync"
)

func max() {
  runtime.GOMAXPROCS(3)
  wg := sync.WaitGroup{}
  wg.Add(20)

  for i := 0; i < 10; i++ {
    go func() {
      fmt.Println("xxx-----------",i)
      wg.Done()
    }()
  }

  for i := 0; i < 10; i++ {
   go func(i int) {
     fmt.Println("yyy-----------",i)
     wg.Done()
   }(i)
  }

  wg.Wait()
}

func main() {
  max()
}
