package ColorfulRabbit

import (
  "fmt"
  "testing"
)

func TestMd5V(t *testing.T) {
  s := "hello"
  fmt.Println(Md5V(s))
}
