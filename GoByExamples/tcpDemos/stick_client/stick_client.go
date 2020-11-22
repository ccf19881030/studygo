package main

import (
	"fmt"
	"net"
)

// TCP 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
	   msg := `Hello, hello, How are you?`
	   conn.Write([]byte(msg))
	}
}
