package datetime

import (
	"testing"
	"time"
)

func TestGetDateTimeOne(t *testing.T) {
	now := time.Now() 
	nowRes := GetDateTimeOne(now)
  t.Log(nowRes)

  past := GetTimePastBySeconds(now, 30)
  t.Log(past)

  before := GetTimeBeforeBySeconds(now, 30)
  t.Log(before)
}


