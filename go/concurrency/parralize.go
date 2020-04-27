package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	DoJobs()
}

func DoJobs() ([]int, error) {
	limit := 10000000
	nums := make([]int, limit)
	for i := 0; i < limit; i++ {
		nums[i] = i
	}

	res := make([]int, len(nums))
	errs := make(chan error, len(res))

	cal := func(i int) {
		var err error
		// 发到与请求idx相同的位置
		res[i], err = doComplicatedThings(i)
		if err != nil {
			errs <- err
		}
	}

	t1 := time.Now()
	parallize(context.TODO(), 10, len(nums), cal)
	fmt.Printf("time elapsed: %f", time.Since(t1).Seconds())
	if len(errs) != 0 {
		return nil, <-errs
	}
	return res, nil
}

func parallize(ctx context.Context, workers int, jobs int, doJob func(idx int)) {
	if jobs < workers {
		workers = jobs
	}
	toProcess := make(chan int, jobs)
	for i := 0; i < jobs; i++ {
		toProcess <- i
	}
	close(toProcess)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for idx := range toProcess {
				select {
				case <-ctx.Done():
					return
				default:
					doJob(idx)
				}
			}
		}()
	}
}
