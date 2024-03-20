package main

import "fmt"

func c3() {
  ch := make(chan int)
  for i := 0; i < 10; i++ {
    go func(j int) {
      fmt.Println("写入channel---------", j)
      ch <- j
    }(i)
  }

  // 方式1
  //for i := 0; i < 10; i++ {
  //  <-ch
  //}

  // 方式2
  i := 0
  for {
    fmt.Println("for循环次数 ---------------", i)
    if i == 10 {
      break
    }

    select {
    case i := <-ch:
      fmt.Println("接收到: ", i)
    }

    i++
  }

  // todo: 如果没有for， 则只是监听了channel的第一个就会退出了
  //select {
  //case i := <-ch:
  //  fmt.Println("接收到: ", i)
  //}
}
