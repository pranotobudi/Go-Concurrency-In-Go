package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			fmt.Println("sending: ", i)
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("receiving val:", val)
	}
}
