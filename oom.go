package main

import (
	"fmt"
	"os"
	"time"
)

var identifier []int
var c int = 0

func run() {
	identifier = append(identifier, 1)
	// fmt.Println("--")
	time.Sleep(10 * time.Millisecond)
	c++

	run()
}

func main() {
	done := make(chan bool)

	fmt.Printf("pid: %d\n", os.Getpid())
	for i := 0; i < 10000; i++ {
		go run()
	}

	<-done // Block forever
}
