package ColorfulRabbit

import (
  "fmt"
  "gopkg.in/fatih/set.v0"
  "testing"
)

func TestTmp(t *testing.T)  {
  //a := []string{"a", "b"}
  //b := []string{"a", "c"}
  a := set.New(set.ThreadSafe)
  a.Add("a")
  a.Add("b")

  b := set.New(set.ThreadSafe)
  b.Add("a")
  b.Add("c")

  c := set.Difference(a, b)
  fmt.Println(c.List())
}
