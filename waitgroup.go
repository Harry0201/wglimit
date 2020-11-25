package wglimit

import (
	"sync"
)

type limitWaitGroup struct {
	wg         sync.WaitGroup
	concurrent chan struct{}
}

// 初始化
func NewWaitGroup(size int32) *limitWaitGroup {
	if size <= 0 {
		size = 1
	}
	return &limitWaitGroup{
		concurrent: make(chan struct{}, size),
	}
}

// 计数加1
func (lwg *limitWaitGroup) Add() {
	lwg.concurrent <- struct{}{}
	lwg.wg.Add(1)
}

// 计数减1
func (lwg *limitWaitGroup) Done() {
	<-lwg.concurrent
	lwg.wg.Done()
}

// 阻塞
func (lwg *limitWaitGroup) Wait() {
	lwg.wg.Wait()
}
