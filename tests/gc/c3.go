package main

import "fmt"

func Fibonacci() func() int {
  a, b := 0, 1
  fmt.Println("a:", a, "  b:", b)

  return func() int {
    a, b = b, a+b   // todo: go的闭包写法， 第一次初始化函数之后， a, b不再重新赋值
    return a
  }
}

func main() {
  f := Fibonacci()

  for i := 0; i < 10; i++ {
    fmt.Printf("Fibonacci: %d\n", f())
  }
}
