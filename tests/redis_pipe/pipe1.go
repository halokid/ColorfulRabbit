package main

import (
  "flag"
  "fmt"
  "github.com/garyburd/redigo/redis"
  "log"
  "time"
)

func send(conn redis.Conn, n int) {
  for i := 0; i < n; i++ {
    if err := conn.Send("SET", "key", i); err != nil {
      log.Fatal(err)
    }
  }
  if err := conn.Flush(); err != nil {
    log.Fatal(err)
  }
}

func receive(conn redis.Conn, n int) {
  for i := 0; i < n; i++ {
    _, err := conn.Receive()
    if err != nil {
      log.Fatal(err)
    }
  }
}

func main() {
  totalRequests := flag.Int("n", 100000, "总共请求")
  flag.Parse()
  conn, err := redis.Dial("tcp", "172.20.71.25:6379",
    redis.DialPassword("7wUZcd0#"), redis.DialDatabase(4))
  if err != nil {
    log.Fatal(err)
  }
  start := time.Now()
  go send(conn, *totalRequests)
  receive(conn, *totalRequests)
  t := time.Since(start)
  fmt.Printf("%d requests completed in %f seconds\n", *totalRequests, float64(t)/float64(time.Second))
  fmt.Printf("%f requests / second\n", float64(*totalRequests)*float64(time.Second)/float64(time.Since(start)))
}