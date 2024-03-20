package main
/**
逃逸场景（什么情况才分配到堆中）
栈空间不足逃逸（空间开辟过大）
build:  go build -gcflags=-m  c2.go
 */

func Slice() {
  //s := make([]int, 1000, 1000)    // todo: 1000太少， 不足以逃逸到堆
  s := make([]int, 10000, 10000)      // todo: 试下10000

  for idx, _ := range s {
    s[idx] = idx
  }
}

func main() {
  Slice()
}
