package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("good")
	}
}
