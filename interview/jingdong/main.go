package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	//p := NewSimplePool(20)
	for i := 0; i < 100; i++ {

	}
}

// 协程池
// 1. 创建 worker上限 20，任务无上限
// 2. 没有任务时不断的销毁 worker

type Task interface {
	Execute(exitC chan struct{})
}

type SimplePool struct {
	tasksC     chan Task
	exitC      chan struct{}
	taskQ      []Task
	curTaskNum int64
	maxTaskNum int64
	mu         *sync.Mutex
	wg         *sync.WaitGroup
}

func NewSimplePool(taskCap int) *SimplePool {
	return &SimplePool{
		tasksC:     make(chan Task, taskCap),
		maxTaskNum: int64(taskCap),
	}
}

func (p *SimplePool) Run() {
	for {
		select {
		case t, ok := <-p.tasksC:
			p.wg.Add(1)
			p.taskQ = append(p.taskQ, t)
			go func() {
				defer p.wg.Done()
				if ok && p.curTaskNum < p.maxTaskNum && len(p.taskQ) > 0 {
					p.mu.Lock()
					t = p.taskQ[0]
					p.taskQ = p.taskQ[1:]
					p.mu.Unlock()

					atomic.AddInt64(&p.curTaskNum, +1)
					exitC := make(chan struct{})
					go t.Execute(exitC)
					<-exitC
					atomic.AddInt64(&p.curTaskNum, -1)
				}
			}()
		case <-p.exitC:
			p.wg.Wait()
			return
		default:
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()
				// 并发读写问题
				if p.curTaskNum < p.maxTaskNum && len(p.taskQ) > 0 {
					p.mu.Lock()
					t := p.taskQ[0]
					p.taskQ = p.taskQ[1:]
					p.mu.Unlock()

					atomic.AddInt64(&p.curTaskNum, +1)
					exitC := make(chan struct{})
					go t.Execute(exitC)
					<-exitC
					atomic.AddInt64(&p.curTaskNum, -1)
				}
			}()
		}
	}
}

func (p *SimplePool) AddTask(t Task) {
	p.tasksC <- t
}

func (p *SimplePool) Stop() {
	p.exitC <- struct{}{}
}

type Pool struct {
	tasks      chan Task
	waitQ      []Task
	stop       chan struct{}
	wg         *sync.WaitGroup
	defaultMax int
	mu         sync.Mutex
}

func NewPool(defaultMax int) *Pool {
	return &Pool{
		tasks:      make(chan Task, defaultMax),
		defaultMax: defaultMax,
	}
}

// 不能控制协程池的数量

//func (p *Pool) Start() {
//	for {
//		select {
//		case <-p.stop:
//			close(p.tasks)
//			p.wg.Wait()
//			return
//		case w, ok := <-p.tasks:
//			if ok {
//				p.mu.Lock()
//				p.waitQ = append(p.waitQ, w)
//				p.mu.Unlock()
//			}
//		default:
//			p.mu.Lock()
//			if len(p.waitQ) != 0 {
//				p.mu.Unlock()
//
//				go func() {
//					defer p.wg.Done()
//
//					for {
//						p.mu.Lock()
//						if len(p.waitQ) != 0 {
//							t := p.waitQ[0]
//							p.waitQ = p.waitQ[1:]
//							go t.Execute()
//						} else {
//							break
//						}
//						p.mu.Unlock()
//					}
//
//				}()
//			}
//		}
//	}
//}

func (p *Pool) AddTask(t Task) {
	p.tasks <- t
}

func (p *Pool) Stop() {
	p.stop <- struct{}{}
}
