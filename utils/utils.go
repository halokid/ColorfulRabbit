package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/halokid/ColorfulRabbit/logger"
)

func MapToString[T string | int | interface{}](m map[string]T) string {
  ms, err := json.Marshal(m)
  if err != nil {
    logger.SugarLogger.Errorf("MaptoString() err -->>> %+v", err)
    return ""
  }
  return string(ms)
}

func RunRootPath() string {
  absPath, err := filepath.Abs(os.Args[0])
  // logger.Logger.Debugf("absPath -->>> %+v", absPath)
  if err != nil {
    // logger.Logger.Errorf("Error getting absolute path -->>> %+v", err)
    return "RunRoorPaht error"
  }
  currentDir := filepath.Dir(absPath)
  // logger.Logger.Debugf("RunRootPath -->>> %+v", currentDir)
  return currentDir
}




