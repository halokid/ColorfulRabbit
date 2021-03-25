package rate_limit

type RateLimitDo interface {
  Limit() bool    // ture: pass, false: reject
}

type RateLimit struct {
  // rolling window
  windowTime      int64
  windowCount     uint

}

func (r *RateLimit) SetRollingWindow(windowKey string, ) {

}
