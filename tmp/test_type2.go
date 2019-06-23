package main

/**
关于范型函数的两个要点总结，以前一直以为 范型函数要匹配好传入的参数数据类型，
然后要定义好返回的数据类型，其实不是这样的， 因为传入的参数 和 返回的参数都
是在 调用 范型程序 的代码块去处理的，范型函数根本不用考虑

1.  传入的参数在 调用处 interface 化
2.  返回的数据必须是  interface 化，然后在调用处获得之后，自己根本已经有的数据类型去做适配
 */

import (
  "fmt"
)

type UserT interface {
  GetUser()
}

type User struct {
  Id    int
  Name  string
}

func (u User)GetUser() User {
  return u
}

func PattAll(m map[interface{}]interface{}) interface{} {
  // 范型函数的另外一个要点就是 输出的类型是  interface{}， 然后输出的数据就留给调用这个范型函数的逻辑代码去处理就好
  // 因为调用这个函数的逻辑代码的代码块是有 原来 m 的所有适配的数据类型的， 范型函数根本不用这个数据类型有无的问题
  fmt.Println(m)
  return m
}


// 上面的逻辑可以理解为是把 User 结构体 type 化

func main() {
  u := User{Id: 99, Name: "okok"}

  // 范型函数的注意力逻辑，应该是在把要传入的参数对象  interfaces 化，适配所有的传入参数类型就可以了
  m := make(map[interface{}]interface{})

  var s string
  s = "hello"
  m[0] = s

  m[1] = u

  //fmt.Println(m)
  PattAll(m)
}








