package wglimit

import (
	"sync"
	"sync/atomic"
)

type limitWaitGroup struct {
	Concurrent chan struct{}
	wg         sync.WaitGroup
	max        int
	sync.RWMutex
	total int32
}

// 初始化
func NewWaitGroup(size int) *limitWaitGroup {
	if size <= 0 {
		size = 1
	}
	return &limitWaitGroup{
		Concurrent: make(chan struct{}, size),
	}
}

// 计数加1
func (lwg *limitWaitGroup) Add() {
	lwg.Concurrent <- struct{}{}
	lwg.wg.Add(1)
	atomic.AddInt32(&lwg.total, 1)
}

// 计数增加固定值
func (lwg *limitWaitGroup) AddWithCount(count int) {
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		lwg.Concurrent <- struct{}{}
		lwg.wg.Add(1)
	}

}

// 计数减1
func (lwg *limitWaitGroup) Done() {
	lwg.Lock()
	defer lwg.Unlock()

	if _, ok := <-lwg.Concurrent; !ok {
		return
	}
	lwg.wg.Done()
	atomic.AddInt32(&lwg.total, -1)
}

// 计数减少固定值
func (lwg *limitWaitGroup) DoneWithCount(count int) {
	if count == 0 {
		return
	}

	lwg.Lock()
	defer lwg.Unlock()

	for i := 0; i < count; i++ {
		<-lwg.Concurrent
		lwg.wg.Done()
		atomic.AddInt32(&lwg.total, -1)
	}
}

// 计数清零
func (lwg *limitWaitGroup) DoneAll() {
	lwg.Lock()
	defer lwg.Unlock()

	lwg.wg.Add(0 - int(lwg.total))

	ok := true
	for ok {
		_, ok = <-lwg.Concurrent
	}
}

// 阻塞
func (lwg *limitWaitGroup) Wait() {
	lwg.wg.Wait()
}
