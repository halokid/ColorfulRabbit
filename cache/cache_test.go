package main

import (
  "testing"
)

var key = "halokid"

func TestHash64(t *testing.T) {
  idx := Hash64(key)
  t.Logf("[]byte(halokid) ---------- %+v", []byte(key))
  t.Logf("idx ------------ %+v", idx)
}

func TestCompute(t *testing.T) {
  s, err := compute(key)
  t.Logf("s, err --------- %+v, %+v", s, err)
}