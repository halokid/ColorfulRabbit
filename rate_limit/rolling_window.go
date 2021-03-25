package rate_limit

import "time"

type RollingWindow struct {

}

func (r *RollingWindow) Limit() bool {
  currTime := time.Now().Unix()
  if LimitQueue == nil {
    LimitQueue = make(map[string][]int64)
  }

  if _, ok = LimitQueue[queueName]; !ok {
    LimitQueue[queueName] = make([]int64, 0)
  }

  // queue is not full
  if uint(len(LimitQueue[queueName])) < count {
    LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
    return true
  }
  // queue is full, take the earliest time item of the slice
  earliestTime := LimitQueue[queueName][0]
  // the earliest visit still in the queue, and the time difference is less than timeWindow
  if currTime - earliestTime <= timeWindow {
    return false
  } else {
    // the earliest is expire, throw it
    LimitQueue[queueName] = LimitQueue[queueName][1:]
    LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
  }
  return true
}

func NewRollingWd() RateLimit {
  return &RollingWindow{}
}




