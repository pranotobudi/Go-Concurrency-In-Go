package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(ch chan<- string) {
		time.Sleep(2 * time.Second)
		ch <- "channel1"
	}(ch1)

	go func(ch chan<- string) {
		time.Sleep(2 * time.Second)
		ch <- "channel2"
	}(ch2)

	for i := 0; i < 3; i++ {
		select {
		case val1 := <-ch1:
			fmt.Println("val1: ", val1)
		case val2 := <-ch2:
			fmt.Println("val2: ", val2)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}

	}

}
