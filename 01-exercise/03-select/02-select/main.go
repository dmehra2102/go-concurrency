package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch
	select {
	case m1 := <-ch:
		fmt.Println(m1)
	case <-time.After(2 * time.Second):
		fmt.Println("Time is out")
	}
}
