package guid

import "testing"

func mustNewV5(t *testing.T, namespace GUID, name []byte) GUID {
  t.Helper()

  g, err := NewV5(namespace, name)
  if err != nil {
    t.Fatal(err)
  }
  return g
}

func mustNewV4(t *testing.T) GUID {
  t.Helper()

  g, err := NewV4()
  if err != nil {
    t.Fatal(err)
  }
  return g
}

func TestGenGuid(t *testing.T) {
  g, err := FromString("00000000-0000-0000-0000-0000000000000")
  t.Logf("g -->>> %+v, err -->>> %+v", g, err) 

  // guid := GUID{
  //   Data1: 1,
  //   Data2: 2,
  //   Data3: 3,
  //   Data4: ,
  // }
  guid := mustNewV4(t)
  gid := guid.String()
  t.Log(gid)

  guid2 := mustNewV4(t)
  gid2 := guid2.String()
  t.Log(gid2)
}