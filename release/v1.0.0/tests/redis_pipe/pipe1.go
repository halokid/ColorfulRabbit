package main

import (
  "flag"
  "fmt"
  "github.com/garyburd/redigo/redis"
  "github.com/spf13/cast"
  "log"
  "time"
)

func send(conn redis.Conn, n int) {
  for i := 0; i < n; i++ {
    if err := conn.Send("SET", "key" + cast.ToString(i), i); err != nil {
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
  conn, err := redis.Dial("tcp", "8.8.8.8:222",
    redis.DialPassword("xxxxx"), redis.DialDatabase(4))
  if err != nil {
    log.Fatal(err)
  }
  start := time.Now()
  go send(conn, *totalRequests)
  //receive(conn, *totalRequests - 1000)
  receive(conn, *totalRequests)

  //v, err := redis.Bytes(conn.Receive())
  //log.Println("v ----------------", v, err)
  //time.Sleep(10 * time.Second)

  t := time.Since(start)
  fmt.Printf("%d requests completed in %f seconds\n", *totalRequests, float64(t)/float64(time.Second))
  fmt.Printf("%f requests / second\n", float64(*totalRequests)*float64(time.Second)/float64(time.Since(start)))
}