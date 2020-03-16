package ColorfulRabbit

import (
  "crypto/md5"
  "encoding/hex"
)

/**
加解密相关
 */

func Md5V(s string) string {
  h := md5.New()
  h.Write([]byte(s))
  return hex.EncodeToString(h.Sum(nil))
}
