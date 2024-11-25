package wait

import (
	"sync"
	"time"
)

// Wait 是一个类似于 sync.WaitGroup 的结构体，可以在指定超时的情况下进行等待
type Wait struct {
	wg sync.WaitGroup
}

// Add 将 delta（可以为负数）添加到 WaitGroup 计数器
func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

// Done 将 WaitGroup 计数器减一
func (w *Wait) Done() {
	w.wg.Done()
}

// Wait 阻塞直到 WaitGroup 计数器为零
func (w *Wait) Wait() {
	w.wg.Wait()
}

// WaitWithTimeout 阻塞直到 WaitGroup 计数器为零或超时
// 如果超时返回 true，否则返回 false
func (w *Wait) WaitWithTimeout(timeout time.Duration) bool {
	c := make(chan bool, 1) // 创建一个带缓冲的通道
	go func() {
		defer close(c) // 确保通道在退出时关闭
		w.wg.Wait()    // 等待 WaitGroup 的计数器归零
		c <- true      // 发送完成信号到通道
	}()
	// 使用 select 等待通道消息或时间超时
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
