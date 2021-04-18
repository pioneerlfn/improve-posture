package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
)

type Cache struct {
	cache *lru.Cache
	queue *lru.Cache
}

var checksumCache *Cache

const (
	checksumCacheSize = 100000
)

func init() {
	cache, err := lru.New(checksumCacheSize)
	if err != nil {
		panic(err)
	}
	queue, err := lru.New(checksumCacheSize)
	if err != nil {
		panic(err)
	}
	checksumCache = &Cache{
		cache: cache,
		queue: queue,
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/compute", Compute0)
	r.Run(":8090")
}

func Compute0(c *gin.Context) {
	key := c.Query("key")
	c.JSON(http.StatusOK, Compute(key))
	return
}

func Compute(key string) string {
	if val, ok := checksumCache.cache.Get(key); ok {
		return val.(string)
	}

	if ready, ok := checksumCache.queue.Get(key); ok {
		<-ready.(chan struct{})
		val, _ := checksumCache.cache.Get(key)
		return val.(string)
	}

	checksumCache.queue.Add(key, make(chan struct{}))

	val := compute(key)

	checksumCache.cache.Add(key, val)
	ready, _ := checksumCache.queue.Get(key)
	close(ready.(chan struct{}))

	return val
}

func compute(key string) string {
	sleep := rand.Intn(10)
	time.Sleep(time.Duration(sleep) * time.Second)

	return fmt.Sprintf("0x%x", md5.Sum([]byte(key)))
}
