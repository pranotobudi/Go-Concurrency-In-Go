package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	// var counterLock sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// counterLock.Lock()
				// counter++
				atomic.AddInt64(&counter, 1)
				// counterLock.Unlock()

			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
