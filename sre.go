package ColorfulRabbit

import "os"

/**
系统运维类
 */


// 判断文件是否存在
func PathExists(path string) bool {
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return false
}







