package ColorfulRabbit

import "testing"

func TestReverseSl(t *testing.T) {
  sl := []interface{}{5, 4, 3, 2, 1}
  slx := ReverseSl(sl)
  t.Log(slx)
}
