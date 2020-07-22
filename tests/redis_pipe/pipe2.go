package main

import (
  "flag"
  "fmt"
  "github.com/garyburd/redigo/redis"
  "log"
  "time"
)

type command struct {
  name          string
  args          []interface{}
  result        chan result
}

type result struct {
  err       error
  value     interface{}
}

type runner struct {
  conn    redis.Conn
  send    chan command
  recv    chan chan result
  stop    chan struct{}
  done   chan struct{}
}

func (r *runner) sender() {
  var flush int
  for {
    select {
    case <-r.stop:          // todo: 信号1触发了这里
      if err := r.conn.Flush(); err != nil {
        log.Fatal(err)
      }
      close(r.recv)         // todo: 信号2
      fmt.Println("FLUSH:", flush)
      return

    case cmd := <-r.send:
      if err := r.conn.Send(cmd.name, cmd.args...); err != nil {
        log.Fatal(err)
      }
      if len(r.send) == 0 || len(r.recv) == cap(r.recv) {
        flush++
        if err := r.conn.Flush(); err != nil {
          log.Fatal(err)
        }
      }
      r.recv <-cmd.result
    }
  }
}

func (r *runner) receiver() {
  for ch := range r.recv {        // todo: 信号2触发这for循环退出
    var result result
    result.value, result.err = r.conn.Receive()
    ch <-result
    if result.err != nil && r.conn.Err() != nil {
      log.Fatal(r.conn.Err())
    }
  }
  close(r.done)       // todo: 信号3
}

func newRunner() *runner {
  conn, err := redis.Dial("tcp", "8.8.8.8:222",
    redis.DialPassword("xxxxxx"), redis.DialDatabase(4))
  if err != nil {
    log.Fatal(err)
  }
  r := &runner{
    conn:     conn,
    send:     make(chan command, 100),
    recv:     make(chan chan result, 100),
    stop:     make(chan struct{}),
    done:     make(chan struct{}),
  }

  go r.sender()
  go r.receiver()
  return r
}

func main() {
  totalRequests := flag.Int("n", 100000, "总共请求")
  flag.Parse()
  r := newRunner()
  start := time.Now()
  args := []interface{}{"a", "b"}
  for i := 0; i < *totalRequests; i++ {
    r.send <-command{name:  "SET", args:  args, result:  make(chan result, 1)}
  }
  close(r.stop)       // close也是发信号给stop chan struct{}, todo: 信号1
  <-r.done              // todo: 等待信号3执行，不然就一直阻塞
  t := time.Since(start)
  fmt.Printf("%d requests completed in %f seconds\n", *totalRequests, float64(t)/float64(time.Second))
  fmt.Printf("%f requests / second\n", float64(*totalRequests)*float64(time.Second)/float64(time.Since(start)))
}





