package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var c = sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) < 1 {
			c.Wait()
		}
		fmt.Println("sharedRsc1:", sharedRsc["src1"])
		c.L.Unlock()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) < 2 {
			c.Wait()
		}
		fmt.Println("sharedRsc2:", sharedRsc["src2"])
		c.L.Unlock()
	}()
	c.L.Lock()
	sharedRsc["src1"] = "foo"
	sharedRsc["src2"] = "foo"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
