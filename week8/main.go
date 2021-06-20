package main

import (
	"log"
	"strconv"

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
	for i := 0; i < 10; i++ {
		RCon.Do("DEL", "key:00000000000"+strconv.Itoa(i))
	}

	for i := 10; i < 100; i++ {
		RCon.Do("DEL", "key:0000000000"+strconv.Itoa(i))
	}
}
