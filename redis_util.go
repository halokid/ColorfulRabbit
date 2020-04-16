package ColorfulRabbit

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

var (
  RdsTyp    int         // 0 for pool, 1 for conn
  RdsPool *redis.Pool
  RdsConn redis.Conn
)

func InitRdsPool(host, port, pwd string, db int) {
  RdsPool = &redis.Pool{
    MaxIdle:        50,
    MaxActive:      10000,
    Dial: func() (redis.Conn, error) {
      conn, err := redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db))
      CheckFatal(err, "----------------- redis pool error")
      return conn, err
    },
  }
  RdsTyp = 0
}

func InitRdsConn(host, port, pwd string, db int) {
  var err error
  RdsConn, err = redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db))
  CheckFatal(err, "----------------- redis conn error")
  RdsTyp = 1
  //return RdsConn, err
}

func getConn() redis.Conn {
  if RdsTyp == 0 {
    return RdsPool.Get()
  }
  return RdsConn
}

func RdsPing() error {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  _, err := redis.String(conn.Do("PING"))
  if err != nil {
    return fmt.Errorf("cannot 'PING' db: %v", err)
  }
  return nil
}

func RdsGet(key string) ([]byte, error) {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  var data []byte
  data, err := redis.Bytes(conn.Do("GET", key))
  if err != nil {
    return data, fmt.Errorf("error getting key %s: %v", key, err)
  }
  return data, err
}

func RdsSet(key string, value []byte) error {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  _, err := conn.Do("SET", key, value)
  if err != nil {
    v := string(value)
    if len(v) > 15 {
      v = v[0:12] + "..."
    }
    return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
  }
  return err
}

func RdsExists(key string) (bool, error) {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  ok, err := redis.Bool(conn.Do("EXISTS", key))
  if err != nil {
    return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
  }
  return ok, err
}

func RdsDelete(key string) error {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  _, err := conn.Do("DEL", key)
  return err
}

func RdsGetKeys(pattern string) ([]string, error) {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  iter := 0
  keys := []string{}
  for {
    arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
    if err != nil {
      return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
    }

    iter, _ = redis.Int(arr[0], nil)
    k, _ := redis.Strings(arr[1], nil)
    keys = append(keys, k...)

    if iter == 0 {
      break
    }
  }

  return keys, nil
}

func RdsIncr(counterKey string) (int, error) {

  //conn := RdsPool.Get()
  conn := getConn()
  defer conn.Close()

  return redis.Int(conn.Do("INCR", counterKey))
}


