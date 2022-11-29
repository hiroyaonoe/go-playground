package main

import (
	"bytes"
	"errors"
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

	//	err = readHeader(conn)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("Header BOPP TEST Accepted")

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatalf(err.Error())
	}
}

func readHeader(r io.Reader) error {
	header := []byte("BCoP TEST\r\n")
	b := make([]byte, len(header))

	_, err := r.Read(b)
	if err != nil {
		return err
	}

	if bytes.Equal(b, header) {
		return nil
	}

	return errors.New("Invalid Header")
}
