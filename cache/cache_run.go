package main

import (
  "github.com/gin-gonic/gin"
  lru "github.com/hashicorp/golang-lru"
  "os"
  "sync"
)

type Cache struct {
  cache         *lru.Cache
  mutexBucket   map[uint64]*sync.RWMutex
}

var cache *Cache

const cacheSize = 100000

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
}

func main() {
  r := gin.Default()
  gin.DefaultWriter = os.Stdout

}



