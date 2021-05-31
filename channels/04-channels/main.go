package main

import "fmt"

func generateMsg(ch1 chan<- string) {
	ch1 <- "generate message ch1"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	val := <-ch1
	ch2 <- val
	close(ch2)
}

func main() {
	//create ch1 ch2
	ch1 := make(chan string)
	ch2 := make(chan string)
	go generateMsg(ch1)
	go relayMsg(ch1, ch2)

	output := <-ch2
	fmt.Println("output: ", output)
}
