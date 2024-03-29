package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建一个有 5 个工作协程的协程池
	pool := NewGoroutinePool(5)
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	pool.Start(ctx)

	// 提交一些任务到协程池
	for i := 0; i < 10; i++ {
		count := i
		pool.Submit(func() {
			fmt.Printf("Running task %d\n", count)
			time.Sleep(2 * time.Second)
		})
	}

	// 关闭协程池，并等待所有任务完成
	pool.Close()
	fmt.Println("All tasks completed.")
}

type Task func()

type GoroutinePool struct {
	taskChan    chan Task
	workerCount int
	wg          sync.WaitGroup
}

func NewGoroutinePool(workerCount int) *GoroutinePool {
	return &GoroutinePool{
		taskChan:    make(chan Task, workerCount),
		workerCount: workerCount,
	}
}

func (g *GoroutinePool) Start(ctx context.Context) {
	for i := 0; i < g.workerCount; i++ {
		g.wg.Add(1)
		go func() {
			defer g.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case task, ok := <-g.taskChan:
					if ok {
						task()
					} else {
						return
					}
				}
			}
		}()
	}
}

func (g *GoroutinePool) Submit(task Task) {
	g.taskChan <- task
}

func (g *GoroutinePool) Close() {
	close(g.taskChan)
	g.wg.Wait()
}
