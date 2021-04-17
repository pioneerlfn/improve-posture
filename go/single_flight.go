package main

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"

	"github.com/gin-gonic/gin"
)

type Cache struct {
	cache       *lru.Cache
	mutexBucket map[uint64]*sync.RWMutex
}

var cache *Cache

const cacheSize = 100000

func init() {
	c, err := lru.New(cacheSize)
	if err != nil {
		panic(err)
	}
	cache = &Cache{
		cache:       c,
		mutexBucket: make(map[uint64]*sync.RWMutex, cacheSize),
	}
	for i := 0; i < cacheSize; i++ {
		m := &sync.RWMutex{}
		cache.mutexBucket[uint64(i)] = m
	}
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

func Compute(con *gin.Context) {
	//获取已备份版本
	key := con.Query("key")

	mutexIdx := Hash64(key) % cacheSize
	cache.mutexBucket[mutexIdx].Lock()
	defer cache.mutexBucket[mutexIdx].Unlock()

	val, ok := cache.cache.Get(key)
	if ok {
		con.JSON(http.StatusOK, val)
		return
	}

	value, err := compute(key)
	if err != nil {
		con.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	updateCache(key, value)
	con.JSON(http.StatusOK, value)
}

func updateCache(key, val string) {
	cache.cache.Add(key, val)
}

func compute(key string) (string, error) {
	sleep := rand.Intn(10)
	fmt.Printf("[key: %s], sleep %ds\n", key, sleep)
	// 随机sleep, 模拟不同key的计算时长
	time.Sleep(time.Duration(sleep) * time.Second)
	return fmt.Sprintf("0x%x", md5.Sum([]byte(key))), nil
}

func Hash64(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}
