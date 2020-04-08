package ColorfulRabbit

import "testing"

func TestGetSpIdx(t *testing.T) {
  s := "tops=0&typ=py"
  gsp := GetSpIdx(s, "&", -1)
  t.Log(gsp)

  gsp = GetSpIdx(s, "&", 1)
  t.Log(gsp)
}
