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

func GetMinBefore(i int) string {
  // 获取当前时间几分钟之前的时间格式
  now := time.Now()
  minBrf := now.Add(time.Duration(-i) * time.Minute)
  minBrfStr := minBrf.Format("2006-01-02 15:04")
  return minBrfStr
}
