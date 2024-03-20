package main

import (
  "fmt"
  "sort"
)

func main() {
  m := make(map[int]int, 10)
  keys := make([]int, 0, 10)
  for i := 0; i < 10; i++ {
    m[i] = i
    keys = append(keys, i)
  }

  fmt.Println("keys --------------", keys)

  // 降序
  //sort.Slice(keys, func(i, j int) bool {
  //  if keys[i] > keys[j] {        // 前一位大于后一位
  //    return true
  //  }
  //  return false
  //})

  // 升序
  sort.Slice(keys, func(i, j int) bool {
    if keys[i] < keys[j] {
      return true
    }
    return false
  })

  for _, v := range keys {
    fmt.Println(m[v])
  }
}
