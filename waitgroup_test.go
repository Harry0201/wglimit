package wglimit

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := NewWaitGroup(20)

	for i := 0; i < 100; i++ {
		wg.Add()
		go func() {
			fmt.Println("Now time is:", time.Now().Unix())
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
}
