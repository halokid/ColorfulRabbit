package tests

import (
  "fmt"
  . "github.com/halokid/ColorfulRabbit"
  "reflect"
  "testing"
)

func TestKeys(t *testing.T) {
  var m Xmap
  m = make(map[interface{}]interface{})
  //m = make(map[int]interface{})

  m[1] = "hello"
  m[2] = "world"
  m[3] = "ok"

  keys := m.Keys()
  fmt.Println(keys)
}

func TestItems(t *testing.T) {
  var m Xmap
  m = make(map[interface{}]interface{})

  m[1] = "hello"
  m[2] = "world"
  m[3] = "ok"

  items := m.Items()
  fmt.Println(items)
}

func TestTmp(t *testing.T) {
  //it := []int{1, 2, 3, 4}
  //retSlc([]int{1, 2, 3, 4})

  it := []interface{}{1, 2, 3, 4}
  retSlc(it)

  var ix interface{}
  fmt.Println(ix)
}

func retSlc(it []interface{}) {
  fmt.Println(it)
}

func TestPatt(t *testing.T) {
  var i int = 9
  patt(i)

  slc := []int{1, 2, 3}
  patt(slc)

  mx := make(map[interface{}]interface{})
  mx[1] = "hello"
  mx[2] = "world"
  pattx(mx)

  x := 9
  xi := intToIter(x)
  val := reflect.ValueOf(xi)
  fmt.Println(val.Kind())

}

func patt(it interface{}) {
  fmt.Println(it)
}

func intToIter(i int) interface{} {
  return i
}

func pattx(iter interface{}) {
  t := reflect.TypeOf(iter)
  fmt.Println(t)

  switch vv := iter.(type) {
  case map[interface{}]interface{}:
    fmt.Println("ok")
    fmt.Println(vv)     // vv 是变量实际的值
  }

  iterx := iter.(map[interface{}]interface{})
  for i, v := range iterx {
    fmt.Println(i, "------", v)
  }
}













