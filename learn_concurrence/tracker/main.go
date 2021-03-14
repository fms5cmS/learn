package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tr := NewTracker()
	// 将并发行为交给调用者！
	// 管控声明周期：当 Tracker 对象中的 ch 被关闭时（调用 Shutdown()），这个 goroutine 就会退出；通过 context 的超时也可以使其退出
	// 而 goroutine 退出时又会向 stop 中发送信号，所以在 Shutdown() 中可以通过 stop 知道何时退出
	go tr.Run()
	_ = tr.Evene(context.Background(), "test")
	_ = tr.Evene(context.Background(), "test")
	_ = tr.Evene(context.Background(), "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker {
	return &Tracker{ch: make(chan string, 10)}
}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func (t *Tracker) Evene(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
