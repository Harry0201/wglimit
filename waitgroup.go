package wglimit

import (
	"sync"
)

type limitWaitGroup struct {
	wg         sync.WaitGroup
	concurrent chan struct{}
}

// NewWaitGroup 初始化
func NewWaitGroup(size int32) *limitWaitGroup {
	if size <= 0 {
		size = 1
	}
	return &limitWaitGroup{
		concurrent: make(chan struct{}, size),
	}
}

// Add 计数加1
func (lwg *limitWaitGroup) Add() {
	lwg.concurrent <- struct{}{}
	lwg.wg.Add(1)
}

// Done 计数减1
func (lwg *limitWaitGroup) Done() {
	<-lwg.concurrent
	lwg.wg.Done()
}

// Wait 阻塞
func (lwg *limitWaitGroup) Wait() {
	lwg.wg.Wait()
}
