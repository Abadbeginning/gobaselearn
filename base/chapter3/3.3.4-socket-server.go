package main

import (
	"fmt"
	"net"
)

func main() {
  net.Listen("tcp", "127.0.0.1:8080")
  fmt.Println("Hello World")
}
