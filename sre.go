package ColorfulRabbit

import (
  "bytes"
  "fmt"
  "io"
  "os"
  "os/exec"
  "strings"
)

/**
系统运维类
 */


func PathExists(path string) bool {
  // 判断文件是否存在
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return false
}


func OsExecOut(cmdStr string)  {
  // 执行系统命令,并且获取输出
  fmt.Println("开始执行命令", cmdStr)
  cmdSpl := strings.Split(cmdStr, " ")
  var cmdSplTr []string
  //cmdSplTr := make([]string, 0)
  for _, v := range cmdSpl {
    //fmt.Println(v, len(v))
    vx := strings.Trim(v, " ")
    //fmt.Println(len(vx))
    if len(v) == 0 {
      continue
    }
    cmdSplTr = append(cmdSplTr, vx)
    //fmt.Println("----------------------")
  }
  //fmt.Println(cmdSplTr)
  cmdArgs := cmdSplTr[1:]
  //fmt.Println(cmdArgs)
  //os.Exit(11)
  cmdx := exec.Command(cmdSplTr[0], cmdArgs...)
  stdout, err := cmdx.StdoutPipe()
  CheckFatal(err, " 执行命令", cmdStr, "输出失败")

  var errbuf bytes.Buffer
  cmdx.Stderr = &errbuf
  err = cmdx.Start()
  CheckError(err)
  _, err = io.Copy(os.Stdout, stdout)
  CheckError(err)
  cmdx.Wait()
  fmt.Println(errbuf.String())
}









