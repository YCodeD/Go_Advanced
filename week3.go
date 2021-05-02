package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	group, ctx := errgroup.WithContext(context.Background())

	cClose := make(chan struct{})      // API 关闭信号
	cSignal := make(chan os.Signal, 0) // signal 关闭信号

	// signal 信号注册
	signal.Notify(cSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)

	// 定义 server API
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world!")
	})
	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		cClose <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// errgroup 启动 http server
	group.Go(func() error {
		return server.ListenAndServe()
	})

	group.Go(func() error {
		select {
		case <-cClose:
			log.Println("API close")
		case <-cSignal:
			log.Println("system close signal")
		case <-ctx.Done():
			log.Println("context close")
		}

		log.Println("server is shutdowning")
		return server.Shutdown(ctx)
	})

	// errgroup 退出
	if err := group.Wait(); err != nil {
		log.Println("errgroup exiting")
	}
}
