package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}
func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
