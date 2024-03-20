package main

import "fmt"

type Duck interface {
  Quack()
}

type Animal struct {
  name string
}

func (animal Animal) Quack() {
  fmt.Println(animal.name, ": Quack! Quack! Like a duck!")
}

func PrMap(m map[int]interface{}) {
  fmt.Println(m)
}


func main() {
  unknownAnimal := Animal{name: "Unknown"}

  var equivalent Duck
  equivalent = unknownAnimal
  equivalent.Quack()

  // 比如你有一个 函数参数形式是  map[int][interface{}
  // 那么你就先新建一个这样的数据结构， 然后把你现在要传进函数的变量装进这个数据结构里面去
  // 类似  m[1] = unknownAnimal
  m := make(map[int]interface{})

  var s string
  s = "hello"
  //m[0] = s.(interface{})
  m[0] = s

  m[1] = unknownAnimal

  fmt.Println(m)

  // 然后这里再调用就可以了， 记得 PrMap 函数是要返回  interface 类型的数据
  PrMap(m)
}






