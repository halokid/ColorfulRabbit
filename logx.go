package ColorfulRabbit
/**
日志类使用函数
 */
import (
  "log"
  "os"
)

type Logx struct {
  DebugFlag  bool
  LogFlFlag    bool
}

// 普通错误检查
func CheckError(err error, output ...interface{}) {
  if err != nil {
    log.Printf("[ERROR] %v --------------- %v\n", err, output)
  }
}

// 致命错误检查
func CheckFatal(err error, output ...interface{}) {
  if err != nil {
    log.Printf("[FATAL] %v -------------- %v\n", err, output)
    os.Exit(500)
  }
}

// 需要输出的调试信息
func (lgx *Logx) DebugPrint(output ...interface{}) {
  if lgx.DebugFlag {
    log.Printf("[DEBUG] ---------------- %v\n", output)
  }

  if lgx.LogFlFlag {
    logFileHandle, err := os.OpenFile("./debug_log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
    CheckError(err)
    sc := ""
    for _, v := range output {
      sc += v.(string)
    }
    _, _ = logFileHandle.Write([]byte(sc))
    logFileHandle.Close()
  }
}



