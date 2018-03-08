package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(string([]byte("Hello Go")))
	time.Sleep(2 * time.Second)
	fmt.Printf("👋🏻 elasped nanoseconds: %d", int(time.Since(now)))
}
