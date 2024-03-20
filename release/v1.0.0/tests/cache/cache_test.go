package main

import (
  "testing"
)

var key = "halokid"

func TestHash64(t *testing.T) {
  idx := Hash64(key)
  t.Logf("[]byte(halokid) ---------- %+v", []byte(key))
  t.Logf("idx ------------ %+v", idx)

  mutexIdx := idx % cacheSize
  t.Logf("mutexIdx ------------ %+v", mutexIdx)

  idxx := Hash64("halokic")
  t.Logf("idxx ------------ %+v", idxx)
  mutexIdxx := idxx % cacheSize
  t.Logf("mutexIdxx ------------ %+v", mutexIdxx)
}

func TestCompute(t *testing.T) {
  s, err := compute(key)
  t.Logf("s, err --------- %+v, %+v", s, err)
}