package main

import (
	"context"
	"fmt"
	"net/http"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

// 1. 注意 goroutine 的生命周期，防止其泄露
// 2. main 如何感知到开启的 goroutine 的退出？两个 goroutine 结束后一定会给 done 发送信号，所以从 done 这个 chan 来感知
// 3. main 如何让 goroutine 退出？通过 close(stop) 来唤醒 goroutine 内部的 Shutdown() 从而让 ListenAndServe() 返回
func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	// 业务进程
	go func() {
		done <- serve("8080", nil, stop)
	}()
	// profiling 进程
	go func() {
		done <- serve("8001", nil, stop)
	}()
	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
