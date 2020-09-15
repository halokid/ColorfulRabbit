package main
/**
逃逸场景（什么情况才分配到堆中）
指针逃逸
 */
type Student struct {
  Name    string
  Age     int
}

func StudentRegister(name string, age int) *Student {
  s := new(Student)   // 局部变量s逃逸到堆

  s.Name = name
  s.Age = age

  return s
}

func main() {
  StudentRegister("Xx", 18)
}
