package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/redis.v5"
)

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return client
}

func main() {
	fmt.Println("Hello world. Go away")
	client := newClient()

	pong, err := client.Ping().Result()
	if err == nil {
		d1 := []byte("hello\ngo\n")
		ioutil.WriteFile("/tmp", d1, 0644)
	}
	fmt.Println(pong, err)

	str, errBr := client.BRPop(0, "Post_Object").Result()

	if errBr != nil {
		fmt.Println(errBr)
	} else {
		fmt.Println(str)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
