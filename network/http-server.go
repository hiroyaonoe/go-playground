package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal(err)
	}

	myln := &MyListener{
		listener: ln,
		count:    0,
	}

	log.Fatal(server.Serve(myln))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func ConnContext(ctx context.Context, c net.Conn) context.Context {
	return context.WithValue(ctx, "conn", c)
}

type MyListener struct {
	listener net.Listener
	count    int
	mu       sync.Mutex
}

var _ net.Listener = &MyListener{}

func (l *MyListener) Accept() (net.Conn, error) {
	conn, err := l.listener.Accept()
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	count := l.count
	fmt.Printf("Conn%d\n", count)
	l.count += 1
	l.mu.Unlock()

	fmt.Printf("Accept\n")

	return &TeeConn{conn}, nil
}

func (l *MyListener) Close() error {
	fmt.Printf("Close\n")
	return l.listener.Close()
}

func (l *MyListener) Addr() net.Addr {
	fmt.Printf("Addr\n")
	return l.listener.Addr()
}

type TeeConn struct {
	net.Conn
}

var _ net.Conn = &TeeConn{}

func (c *TeeConn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)
	fmt.Printf(string(b))
	return
}
