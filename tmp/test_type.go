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

func main() {
  unknownAnimal := Animal{name: "Unknown"}

  var equivalent Duck
  equivalent = unknownAnimal
  equivalent.Quack()

  m := make(map[int]interface{})

  var s string
  s = "hello"
  //m[0] = s.(interface{})
  m[0] = s

  m[1] = unknownAnimal

  fmt.Println(m)
}






