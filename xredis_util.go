package ColorfulRabbit

/**
redis struct 版本
 */
import (
  "fmt"
  "github.com/garyburd/redigo/redis"
  "log"
  "time"
)

type XRedis struct {
  Rds       redis.Conn
  RdsPool   *redis.Pool
  CloseNotAtOnece  bool
}

func NewXR(host, port, pwd string, db int) (*XRedis, error) {
  //Rds, err := redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db))
  Rds, err := redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db), redis.DialConnectTimeout(2 * time.Second))
  //Rds, err := redis.DialTimeout("tcp", host + ":" + port, redis.DialPassword(pwd), redis.DialDatabase(db))
  CheckError(err, "redis newconn err")
  return &XRedis{ Rds: Rds}, err
}

func NewXrPool(host, port, pwd string, db int) (*XRedis, error) {
  rdsPool := &redis.Pool{
    MaxIdle:          20,
    MaxActive:        7000,
    IdleTimeout:      60 * time.Second,
    Wait:             true,
    Dial: func() (redis.Conn, error) {
      conn, err := redis.Dial("tcp", host + ":" + port, redis.DialPassword(pwd),
        redis.DialDatabase(db), redis.DialConnectTimeout(3 * time.Second))
      if err != nil {
        CheckFatal(err, "redis服务不可用")
        return nil, err
      }
      return conn, err
    },
  }
  return &XRedis{RdsPool:  rdsPool}, nil
}

func (x *XRedis) Close() error {
  // todo: 这里执行的是  x.RdsPool.Close() 方法, 因为 GetConn 返回的必然是一个 x.RdsPool.Get()， 在初始化 rdsPool 的时候， 根本就没有设置 x.Rds 的数据
  x.Rds.Close()
  return nil
}

func (x *XRedis) GetConn() redis.Conn {
  if x.Rds != nil {     // todo: two case here, when we use redisPoll, x.Rds is nil, the NewXrPool or NewXR decide this
    return x.Rds
  }
  return x.RdsPool.Get()
}

// todo: just get one key
func (x *XRedis) GetKey(pattern string) ([]string, error) {
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()      // todo: if RedisPool, then put back to pool,  if redis.Conn, then close the conn

  iter := 0
  var keys []string
  for {
    //arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
    arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern, "COUNT", 5000))
    if err != nil {
      return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
    }

    //log.Printf("arr 0 ------------ %+v", string(arr[0].([]uint8)))
    iter, _ = redis.Int(arr[0], nil)
    k, _ := redis.Strings(arr[1], nil)
    //log.Printf("k ------------ %+v", k)
    keys = append(keys, k...)

    if iter == 0 {      // iter为游标， 为0则代表遍历结束
      break
    }

    if len(k) > 0 {       // 当命中一个key，退出
      break
    }
  }

  return keys, nil
}

// todo: get all match keys
func (x *XRedis) GetKeys(pattern string) ([]string, error) {
  conn := x.GetConn()
  //conn := x.Rds
  defer conn.Close()

  iter := 0
  var keys []string
  for {
    //arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
    arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern, "COUNT", 5000))
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


func (x *XRedis) XKeyExist(pattern string) ([]string, error) {
  // scan判断key是否存在
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()

  iter := 0
  var keys []string
  for {
    //arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
    arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern, "COUNT", 5000))
    if err != nil {
      return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
    }

    iter, _ = redis.Int(arr[0], nil)
    log.Println("xKeyExist iter ------------------ ", iter)
    k, _ := redis.Strings(arr[1], nil)
    log.Println("xKeyExist k ------------------ ", k)
    //if len(k) > 0 {
    //  os.Exit(1)
    //}
    keys = append(keys, k...)

    if iter == 0 {
      break
    }
  }

  return keys, nil
}



func (x *XRedis) HGetAll(key string, field ...string) (map[string]interface{}, error) {
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()

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

func (x *XRedis) HGet(key string, field string) (string, error) {
  //conn := x.Rds
  conn := x.GetConn()
  if !x.CloseNotAtOnece {
    defer conn.Close()
  }

  res, err := redis.String(conn.Do("HGET", key, field))
  //log.Println("res -------------", res, err)
  return res, err
}

func (x *XRedis) Get(key string) ([]byte, error) {
  //conn := x.Rds
  conn := x.GetConn()
  if !x.CloseNotAtOnece {
    defer conn.Close()
  }

  var data []byte
  data, err := redis.Bytes(conn.Do("GET", key))
  if err != nil {
    return data, fmt.Errorf("error get key %s: %v", key, err)
  }
  return data, err
}

func (x *XRedis) Set(key string, val string) ([]byte, error) {
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()

  var data []byte
  data, err := redis.Bytes(conn.Do("SET", key, val))
  if err != nil {
    return data, fmt.Errorf("error set key %s: %v", key, err)
  }
  return data, err
}

func (x *XRedis) HSetAll(key string, m map[string]interface{}) error {
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()

  _, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(m)...)
  CheckError(err, "redis HSetAll error")
  return err
}

func (x *XRedis) GetRangeKeys(pattern string, start, end int) {

  conn := x.GetConn()
  //conn := x.Rds
  defer conn.Close()

  //iter := 0
  tmp := start
  var keys []string
  for {
    //arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
    arr, err := redis.Values(conn.Do("SCAN", start, "COUNT", end, "MATCH", pattern))
    if err != nil {
      log.Println(err)
      //return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
    }

    start, _ = redis.Int(arr[0], nil)
    log.Println("start ----------------", start)
    k, _ := redis.Strings(arr[1], nil)
    keys = append(keys, k...)

    if start != tmp {
      break
    }
  }
  //log.Println("len keys ----------------", len(keys))
  //log.Println("keys ----------------", keys)

  //return keys, nil
}

func (x *XRedis) PipeGet(keys []string) [][]byte {
  // piple方式批量get
  //conn := x.Rds
  conn := x.GetConn()
  defer conn.Close()

  var bs [][]byte

  for _, k := range keys {
    err := conn.Send("GET", k)
    CheckError(err)
  }
  conn.Flush()
  for i := 0; i < len(keys); i++ {
    v, err :=conn.Receive()
    CheckError(err)
    //log.Printf("v --------- %+v", string(v.([]byte)))
    bs = append(bs, v.([]byte))
  }
  return bs
}

func (x *XRedis) DelKeys(keys []string) error {
  conn := x.GetConn()
  defer conn.Close()
  for _, key := range keys {
    res, err := conn.Do("DEL", key)
    CheckError(err, "delete key", key, "fail error")
    log.Println("DelKeys ---", key, ", res ---", res)
  }
  return nil
}




