package main

import (
	v1 "book_project/api/book/v1"
	"book_project/internal/service"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	service := service.NewBookService()
	s := grpc.NewServer()
	v1.RegisterBookServerServer(s, service)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(errors.Wrap(err, "start server port :8080"))
	}

	s.Serve(l)
}
