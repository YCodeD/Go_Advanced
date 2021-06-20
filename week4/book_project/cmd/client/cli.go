package main

import (
	v1 "book_project/api/book/v1"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	cli := v1.NewBookServerClient(conn)

	res, err := cli.GetBookInfo(context.Background(), &v1.GetBookInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
