package slice

import "testing"

func Test_InSlice(t *testing.T) {
	a := 9
	as := []int{1, 3, 9, 12}
  in := InSlice(a, as)
  t.Logf("%+v", in)

  b := 11
  in = InSlice(b, as)
  t.Logf("%+v", in)

  c := "abc" 
	bs := []string{"abc", "efg"}
  in = InSlice(c, bs)
  t.Logf("%+v", in)

  d := "ooh" 
  in = InSlice(d, bs)
  t.Logf("%+v", in)
}