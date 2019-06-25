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

func pattAll(iter ...interface{}) {
  //fmt.Println(iter)

  //t := reflect.TypeOf(iter)
  //fmt.Println(t)

  //vx := iter.(map[interface{}]interface{})
  //vx := iter

  //for i, v := range vx {
  //for i, v := range iter {
  //  fmt.Println(i, "--------", v)
  //}

  for _, x := range iter {
    switch xx := x.(type) {
    case string:
      fmt.Println(xx)
    case map[int]string:
      fmt.Println("yyyy")
    default:
      fmt.Println("default deal")
    }
  }
}

func TestPattAll(t *testing.T) {
  m := make(map[int]string)

  m[1] = "hello"
  m[2] = "world"

  pattAll(m)
}


type User struct {
  Id      int
  Name    string
}


type UserT interface {
  GetUser()
}

func (u User)GetUser() User {
  //u := User{Id: 88,  Name: "yy"}
  return u
}


func pattyy(i interface{})  {
  //t := reflect.TypeOf(i).Implements(i)
  //fmt.Println(t)

  ix := i.(interface{})
  fmt.Println(reflect.TypeOf(ix))

  mx := make(map[int]interface{})

  mx[0] = ix
  fmt.Println(mx)


  u := User{Id:  99,  Name: "xx"}
  mu := make(map[int]User)

  mu[0] = u

  // 这里会提示 invalid type assertion: i1.(T1) (non-interface type *T1 on left)
  //ux := u.(interface{})

  // 这里会提示 invalid type assertion: i1.(T1) (non-interface type *T1 on left)
  //ux := u.(User)

  var ui UserT
  uix := ui.(interface{})
  mx[1] = uix
  fmt.Println(mx)
}

func TestPattyy(t *testing.T) {
  i := 10
  pattyy(i)
}













