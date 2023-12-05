package main

import (
	"fmt"
	"log"
	"net"

	bcopnet "github.com/hiroyaonoe/bcop-go/protocol/net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	bconn := bcopnet.ReceiverConn(conn)

	var co net.Conn = bconn

	if tc, ok := co.(*net.TCPConn); ok {
		fmt.Println("ok: cast TCPConn")
		if err := tc.SetKeepAlive(true); err != nil {
			fmt.Println("failed: set keepalive")
		}
	}
	fmt.Println("end")
}
