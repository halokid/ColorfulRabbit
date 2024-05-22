package yuque

import (
  "testing"
)

func TestYuque_GenApi(t *testing.T) {
  y := NewYuque()
  getDocApi := y.GenApi("getDoc", "kb", "rmgv7k")
  t.Log("getDocApi ------------", getDocApi)
}

func TestYuque_DoGet(t *testing.T) {
  y := NewYuque()
  getDocApi := y.GenApi("getDoc", "kb", "rmgv7k")
  t.Log("getDocApi ------------", getDocApi)
  rsp := y.DoGet(getDocApi)
  t.Log("rsp -------------", rsp)
}

func TestYuque_GetDoc(t *testing.T) {
  y := NewYuque()
  bodyHtml := y.GetDoc("kb", "rmgv7k")
  t.Log("bodyHtml ---------", bodyHtml)
}