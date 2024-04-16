package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func MapToString[T string | int | interface{}](m map[string]T) string {
  ms, err := json.Marshal(m)
  if err != nil {
    return ""
  }
  return string(ms)
}

func RunRootPath() string {
  absPath, err := filepath.Abs(os.Args[0])
  if err != nil {
    return "RunRoorPaht error"
  }
  currentDir := filepath.Dir(absPath)
  return currentDir
}

func RemoveEmptyElements(s []string) []string {
  result := []string{}
  for _, str := range s {
    if str != "" || strings.Replace(str, " ", "", -1) != "" {
      result = append(result, str)
    }
  }
  return result
}






