package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	connect, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// some event happens
	i := 0
	for {
		// heartbeat
		connect.Write([]byte("HIHI"))
		fmt.Println("Send Data : ", "HIHI", i)
		i++
		time.Sleep(time.Second * 1)
	}
}
