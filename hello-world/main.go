package main

import (
	"fmt"
	"sync"
	"time"
)

func fun(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(1)
	}
}

func funcCall(wg *sync.WaitGroup, msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// fun("bismillah")
	// wg.Add(1)
	// go funcCall(&wg, "bismillah")
	// wg.Wait()

	wg.Add(1)
	go func(wg *sync.WaitGroup, msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, i)
		}
		wg.Done()
	}(&wg, "bismillah")
	wg.Wait()
}
