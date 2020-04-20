package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

// 使用worker模拟锁的抢占
func worker(key string) error {
	endpoints := []string{"127.0.0.1:2379"}

	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Println("new cli error:", err)
		return err
	}

	sess, err := concurrency.NewSession(cli)
	if err != nil {
		return err
	}

	m := concurrency.NewMutex(sess, "/"+key)

	err = m.Lock(context.TODO())
	if err != nil {
		log.Println("lock error:", err)
		return err
	}

	defer func() {
		err = m.Unlock(context.TODO())
		if err != nil {
			log.Println("unlock error:", err)
		}
	}()

	/*log.Println("get lock: ", n)
	n++*/
	time.Sleep(time.Second) // 模拟执行代码

	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			err := worker("lockname")
			if err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
}
