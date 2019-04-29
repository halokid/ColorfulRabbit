package ColorfulRabbit
/**
数学算法函数
 */
import (
  "fmt"
  "math/rand"
  "strconv"
  "time"
)

// 容量大小， kb转gb
func KbToGbInt(kb int64) int64 {
  gb := kb / 1024 / 1024
  return gb
}


// 容量大小， kb转gb
func KbToGb(kb int64) float64 {
  kbI := float64(kb)
  kbx := kbI / 1024 / 1024
  kbxS := fmt.Sprintf("%.1f", kbx)
  kbxSF, _ := strconv.ParseFloat(kbxS, 64)
  return kbxSF
}


// 获取区间随机数
func RandInt(min, max int) int {
  rand.Seed(time.Now().Unix())
  randNum := rand.Intn(max - min) + min
  return randNum
}



