package main

import (
	"fmt"
	"sync"
)

func load() {
	fmt.Println("should be printed once")
}
func main() {
	var wg sync.WaitGroup
	var o sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.Do(load)
		}()
	}
	wg.Wait()
}
