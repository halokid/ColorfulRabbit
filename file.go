package ColorfulRabbit

import (
  "io/ioutil"
  "regexp"
  "strings"
)

func ReplCtx(file string, old string, new string) error {
  // 代替文件的部分内容
  read, err := ioutil.ReadFile(file)
  CheckFatal(err)
  newCtx := strings.Replace(string(read), old, new, -1)
  err = ioutil.WriteFile(file, []byte(newCtx), 0)
  CheckFatal(err)
  return err
}


func GetMatCtx(file string, regex string) [][]string {
  // 批量获取符合规则的文件里的内容
  r, err := ioutil.ReadFile(file)
  if err != nil {
    CheckFatal(err)
  }
  rege := regexp.MustCompile(regex)
  ps := rege.FindAllStringSubmatch(string(r), -1)
  return ps
}



