package main

import "fmt"

func main() {
  a := []int{1, 2, 3}
  b := a[:2]
  fmt.Println("a ---------", a)
  fmt.Println("b ---------", b)

  a[0] = 10
  // 此时a和b共享底层数组
  fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
  fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
  fmt.Println("----------------------------------")

  b = append(b, 999)
  fmt.Println("a ---------", a)
  fmt.Println("b ---------", b)
  fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
  fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
  fmt.Println("----------------------------------")

  // 同上， 没重新分配， 所以 a[2], b[2] 都被修改
  a[2] = 555
  fmt.Println("a ---------", a)
  fmt.Println("b ---------", b)
  fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
  fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
  fmt.Println("----------------------------------")

  // 超出了cap， 这时b进行重新分配, b[3]=777,cap(b)=6
  // todo: b重新分配内存之后， 再修改b就不会影响a了
  b = append(b, 777)
  a[2] = 666
  fmt.Println("a ---------", a)
  fmt.Println("b ---------", b)
  fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
  fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
}
