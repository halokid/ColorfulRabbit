package ColorfulRabbit

import "testing"

func TestReverseSl(t *testing.T) {
  sl := []interface{}{5, 4, 3, 2, 1}
  slx := ReverseSl(sl)
  t.Log(slx)
}

func TestComm(t *testing.T) {
  s1 := []string{"a", "b", "c", "d", "e"}
  s2 := []string{"a", "f", "k", "d", "b"}

  var s3 []string
  s3 = append(s1, s2...)
  t.Log(s3)
  s4 := RmRepeatSlcStr(s3)
  t.Log(s4)
}