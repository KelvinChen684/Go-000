package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return serve(ctx, ":8080", http.HandlerFunc(serveApp))	// 启动serveApp服务
	})
	eg.Go(func() error {
		return serve(ctx, ":8081", http.HandlerFunc(serveDebug))	// 启动serveDebug服务
	})
	eg.Go(func() error {
		return sig(cancel)	// 信号注册与接收
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("quit by error: %s\n", err)
	}
}

func serve(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
	}

	// 优雅退出
	go func() {
		select {
		case <- ctx.Done():		// 阻塞, 接收退出信号
			s.Shutdown(context.Background())
		}
	}()

	return s.ListenAndServe()
}
func serveApp(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "I love go forever!\n")
}

func serveDebug(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "No, You don't, unless you keep programing with go until get success!\n")
}

func sig(cancel context.CancelFunc) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s := <- c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		cancel()
		fmt.Printf("server quit normally by a quit signal: %s\n", s.String())
		return nil	// 为保证优雅退出，此处必须为nil
	}
	return errors.New("Errors from signal listen!")
}
