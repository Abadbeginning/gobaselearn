package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 任务处理器
type TaskHandler func(interface{})

// 定义任务结构体
type Task struct {
	handler TaskHandler
	param   interface{}
}

// 协程池接口
type WorkerPoolImpl interface {
	// 增加worker
	AddWorker()

	// 发送任务
	SendTask(Task)

	// 释放
	Release()
}

// 协程池
type WorkerPool struct {
	wg   sync.WaitGroup
	inCh chan Task
}

// 添加worker
func (wp *WorkerPool) AddWorker() {
	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		for task := range wp.inCh {
			task.handler(task.param)
		}
	}()
}

// 释放
func (wp *WorkerPool) Release() {
	close(wp.inCh)
	wp.wg.Wait()
}

// 发送任务
func (wp *WorkerPool) SendTask(task Task) {
	wp.inCh <- task
}

// 初始化协程池
func NewWorkerPool(num int) WorkerPoolImpl {
	wp := &WorkerPool{
		inCh: make(chan Task, num),
	}
	return wp
}

func main() {
	// 设置缓冲大小
	bufferSize := 100

	// 初始化协程池
	wp := NewWorkerPool(bufferSize)
	workers := 4

	for i := 0; i < workers; i++ {
		wp.AddWorker()
	}

	var sum int32
	testFunc := func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
	}

	var i, n int32
	n = 100
	for ; i < n; i++ {
		task := Task{
			testFunc,
			i,
		}
		wp.SendTask(task)
	}

	wp.Release()
	fmt.Println(sum)
}
