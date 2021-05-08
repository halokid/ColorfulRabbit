package main
/*
缓存分段锁， 解决大并发情况下要创建大量缓存，在很多请求同时写缓存的情况下， 缓存本身是一个map， 所以锁住， 假如用统一把锁锁住一个大map的话， 不断的lock、unlock，这样会性能很低， 所以我们就创建了很多单独的锁，为每个key的缓存写入的时候，用单独的锁去锁住，这样就会提高性能

缓存分段锁解决的主要问题是， 当大并发情况下， 大量请求同时请求同一个缓存key的情况下， 就会造成重复去这个缓存key的数据源取数据的问题， 这样的话， 只要其中一个请求A先lock， 然后去取数据并且写入缓存，然后该处理unlock 其他的同一个key的请求获取到A请求unlock的信号， 直接去读缓存就可以了
 */
import (
  "crypto/md5"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/hashicorp/golang-lru"
  "hash/fnv"
  "math/rand"
  "net/http"
  "os"
  "sync"
)

type Cache struct {
  cache         *lru.Cache
  mutexBucket   map[uint64]*sync.RWMutex
}

var cache *Cache

const cacheSize = 100000

var mapx map[string]string

func init() {
  c, err := lru.New(cacheSize)
  if err != nil {
    panic(err)
  }
  cache = &Cache{
    cache:         c,
    mutexBucket:   make(map[uint64]*sync.RWMutex, cacheSize),
  }

  for i := 0; i < cacheSize; i++ {
    m := &sync.RWMutex{}
    cache.mutexBucket[uint64(i)] = m
  }

  mapx = make(map[string]string)
}

func main() {
  r := gin.Default()
  gin.DefaultWriter = os.Stdout
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong")
  })
  r.GET("/compute", Compute)
  r.Run(":8080")
}

func Compute(c *gin.Context) {
  // 获取已备份版本
  key := c.Query("key")
  //mapx["key"] = "halokid"     // todo: 大并发同时写map，这里会出错，我的测试是200并发

  // todo: Hash64 计算出的key， 然后取余值作为分段锁, 有可能某些key计算出的 mutexIdx 是一样的
  mutexIdx := Hash64(key) % cacheSize
  cache.mutexBucket[mutexIdx].Lock()      // todo:  这个锁其实是锁了下面整个代码逻辑， 并不是特指锁哪一个缓存赋值的过程
  defer cache.mutexBucket[mutexIdx].Unlock()

  // 假如命中缓存
  val, ok := cache.cache.Get(key)
  if ok {     // 如果命中key
    c.JSON(http.StatusOK, val)
    return
  }

  value, err := compute(key)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }

  updateCache(key, value)
}

// 根据key计算出一串hash值，根据这串hash值计算出一个数字, 只要key不同，这串数字就不会相同
func Hash64(key string) uint64 {
  h := fnv.New64a()
  h.Write([]byte(key))
  return h.Sum64()
}

// 用key计算缓存的值
func compute(key string) (string, error) {
  sleep := rand.Intn(10)
  fmt.Printf("[key: %s], sleep %d\n", key, sleep)
  // 随机sleep， 模拟不同key的计算时长
  //time.Sleep(time.Duration(sleep) * time.Second)
  return fmt.Sprintf("0x%x", md5.Sum([]byte(key))), nil
}

func updateCache(key, val string) {
  cache.cache.Add(key, val)
}




