package main

import "fmt"

func main() {
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 6; i++ {
				fmt.Println("sending:", i)
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		for val := range ch {
			fmt.Println("receive:", val)
		}
	}
	ch := owner()
	consumer(ch)
}
