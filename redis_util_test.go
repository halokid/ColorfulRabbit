package ColorfulRabbit

import (
  "fmt"
  "github.com/wxnacy/wgo/arrays"
  "testing"
)

func TestHasItem(t *testing.T) {
  var arr =  []string{"hello", "world"}
  i := arrays.Contains(arr, "hello")
  fmt.Println(i)
  j := arrays.Contains(arr, "xxxx")
  fmt.Println(j)
}