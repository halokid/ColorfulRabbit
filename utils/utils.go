package until

import (
	"ColorfulRabbit/logger"
	"encoding/json"
)

func MapToString[T string | int | interface{}](m map[string]T) string {
  ms, err := json.Marshal(m)
  if err != nil {
    logger.SugarLogger.Errorf("MaptoString() err -->>> %+v", err)
    return ""
  }
  return string(ms)
}




