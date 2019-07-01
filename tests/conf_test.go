package tests

import (
  "fmt"
  "github.com/Unknwon/goconfig"
  . "github.com/halokid/ColorfulRabbit"
  "testing"
)

func TestConf(t *testing.T) {
  url := GetValx("./test_conf.ini", "jvm", "url")
  fmt.Println(url)

  c, _ := goconfig.LoadConfigFile("./test_conf.ini")
  secs := c.GetSectionList()
  fmt.Println(secs)
}
