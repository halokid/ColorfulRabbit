package ColorfulRabbit

import "time"

/**
  日期时间类
 */

func GetDateHour() string {
  // get date, year:month:day:hour like 2019060115, 15 is hour
  datehStr := time.Now().Format("2006010215")
  return datehStr
}

func GetNowMin() int {
  // get minute now
  t := time.Now()
  return t.Minute()
}
