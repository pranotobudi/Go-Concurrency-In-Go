package main

import "fmt"

func main() {
	ch := make(chan int, 6)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Println("sending: ", i)
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println("receiving: ", val)
	}
}
