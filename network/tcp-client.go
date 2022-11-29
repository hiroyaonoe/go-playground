package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	target, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, target)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "BCoP TEST\r\n")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
