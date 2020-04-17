package ColorfulRabbit

/**
redis struct 版本
 */
import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

type XRedis struct {
  Rds       redis.Conn
  RdsPool   redis.Pool
}

func NewXR(host, port, pwd string, db int) (*XRedis, error) {
  Rds, err := redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db))
  CheckError(err, "redis newconn err")
  return &XRedis{ Rds: Rds}, nil
}

func (x *XRedis) Close() error {
  x.Rds.Close()
  return nil
}

func (x *XRedis) GetKeys(pattern string) ([]string, error) {
  conn := x.Rds
  //defer conn.Close()

  iter := 0
  var keys []string
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


func (x *XRedis) HGetAll(key string, field ...string) (map[string]interface{}, error) {
  conn := x.Rds
  //defer conn.Close()
  keys, err := redis.Values(conn.Do("HKEYS", key))
  CheckError(err, "redis hmget error")
  //return keys, err
  vals, err := redis.Values(conn.Do("HVALS", key))

  hmAll := make(map[string]interface{})
  for i, key := range keys {
    hmAll[string(key.([]uint8))] = string(vals[i].([]uint8))
  }
  return hmAll, nil
}



