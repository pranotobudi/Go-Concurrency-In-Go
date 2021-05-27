package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	// router.GET("/", handleRoot)
	// router.Run(":8080")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	counter := 0
	for {
		counter++
		_, err := io.WriteString(c, fmt.Sprintf("response from server %d \n", counter))
		if err != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func handleRoot(c *gin.Context) {
	for {
		fmt.Println("server response")
		time.Sleep(100 * time.Millisecond)
	}
}
