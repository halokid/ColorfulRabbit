package tests

import (
  "github.com/halokid/ColorfulRabbit"
  "testing"
)

func TestStrToJson(t *testing.T) {
  s := `{"name": "kyle", "age": 18}`
  sJs := ColorfulRabbit.StrToJson()
}
