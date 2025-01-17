package main

import (
	"fmt"
	"time"
)

func worker(done chan string) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- "false"
}

func main() {

	done := make(chan string, 2)
	go worker(done)

	<-done
}
