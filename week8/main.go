package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var RCon redis.Conn

func init() {
	var err error
	RCon, err = redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal("Failed to connect redis: " + err.Error())
	}
}

func main() {

}
