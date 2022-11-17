package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := "localhost:8080"
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	conn, err := list.Accept()
	defer conn.Close()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatalf(err.Error())
	}
}
