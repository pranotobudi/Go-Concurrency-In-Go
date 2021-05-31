package main

import (
	"fmt"
	"sync"
)

var balance int
var wg sync.WaitGroup
var balanceLock sync.Mutex

func deposit(i int) {
	balanceLock.Lock()
	defer balanceLock.Unlock()
	balance++
}
func withdraw(i int) {
	balanceLock.Lock()
	defer balanceLock.Unlock()
	balance--
}
func main() {
	// runtime.GOMAXPROCS(2)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Done()
			deposit(i)
		}()
	}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Done()
			withdraw(i)
		}()
	}
	wg.Wait()
	fmt.Println("balance:", balance)
}
