package error

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "testing"

  "github.com/pkg/errors"
)

// todo: return the error interface
func doSomething() interface{} {
  return &ColorError{"Some Error", "error_test.go", 7}
}

// ------------------------------------------------------

var MyError = errors.New("myError")

// todo: return the special error
func retMyError() error {
  return MyError
}

func TestComm(t *testing.T) {
  err := doSomething()

  // todo: 1. 根据错误类型判断错误, 这种判断方式比较复杂， 这种方式不太行
  switch errx := err.(type) {
  case nil:
    t.Log("success...")
  case *ColorError:
    t.Log("error occurred on line", errx.Line)
  default:
    t.Log("unknow type of the error")
  }

  // todo: 2. 根据错误具体数据判断错误, error interface 的 Error 方法的输出，是给人看的，不是给机器看的。我们通常会把Error方法返回的字符串打印到日志中，或者显示在控制台上。永远不要通过判断Error方法返回的字符串是否包含特定字符串，来决定错误处理的方式。 判断错误类型的方式效率太低， 这种方式不行
  retErr := retMyError()
  t.Log("check error is the same:", retErr==MyError)
  // todo: 判断这种错误类型的方法类似： strings.Contains(error.Error(), "not found") 这样， 好挫
}

// todo: ----------- 一种比较好的， 可以从error的发生 trace 到调用代码上下文的方式 ------------

func ReadFile(path string) ([]byte, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, errors.Wrap(err, "[ERROR]trace到调用1 ---- open failed")
  }
  defer f.Close()
  buf, err := ioutil.ReadAll(f)
  if err != nil {
    return nil, errors.Wrap(err, "read failed")
  }
  return buf, nil
}

func ReadConfig() ([]byte, error) {
  home := os.Getenv("HOME")
  config, err := ReadFile(filepath.Join(home, ".settings.xml"))
  return config, errors.Wrap(err, "[ERROR]错误实际发生处 --- could not read config")
}

func TestComm2(t *testing.T) {
  _, err := ReadConfig()
  if err != nil {
    fmt.Printf("%+v\n", err)
    os.Exit(1)
  }
}

/*
// todo: 上面的报错信息为
open /Users/caichengyang/.settings.xml: no such file or directory
[ERROR]trace到调用1 ---- open failed
main.ReadFile
        /Users/caichengyang/go/src/github.com/ethancai/goErrorHandlingSample/sample4/main.go:15
main.ReadConfig
        /Users/caichengyang/go/src/github.com/ethancai/goErrorHandlingSample/sample4/main.go:27
main.main
        /Users/caichengyang/go/src/github.com/ethancai/goErrorHandlingSample/sample4/main.go:32
runtime.main
        /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go:195
runtime.goexit
        /usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
[ERROR]错误实际发生处 --- could not read config
main.ReadConfig
        /Users/caichengyang/go/src/github.com/ethancai/goErrorHandlingSample/sample4/main.go:28
main.main
        /Users/caichengyang/go/src/github.com/ethancai/goErrorHandlingSample/sample4/main.go:32
runtime.main
        /usr/local/Cellar/go/1.9.2/libexec/src/runtime/proc.go:195
runtime.goexit
        /usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
exit status 1

*/









