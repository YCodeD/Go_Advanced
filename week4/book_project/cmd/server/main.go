package main

import (
	v1 "book_project/api/book/v1"
	"book_project/internal/service"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	service := service.NewBookService()
	s := grpc.NewServer()
	v1.RegisterBookServerServer(s, service)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		log.Println("server shutting down...")
		s.GracefulStop()
		return nil
	})

	g.Go(func() error {
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			return errors.Wrap(err, "start server port :8080")
		}
		fmt.Println("grpc server on :8080")
		return s.Serve(l)
	})

	g.Go(func() error {
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-sign:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	log.Printf("errgroup exiting: %+v\n", g.Wait())

}
