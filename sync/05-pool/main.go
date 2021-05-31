package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("new allocation bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, msg string) {
	// var b bytes.Buffer
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.Write([]byte(msg))
	w.Write(b.Bytes())
	bufPool.Put(b)
}
func main() {
	log(os.Stdout, "bismillah1\n")
	log(os.Stdout, "bismillah2\n")
}
