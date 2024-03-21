package until

import "testing"

func TestMapToString(t *testing.T) {
  m := make(map[string]interface{})
  m["name"] = "John"
  m["age"] = 18
  ms := MapToString(m)
  t.Logf("ms -->>> %+v", ms)
}