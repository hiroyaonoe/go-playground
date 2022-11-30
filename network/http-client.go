package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	BCoPKey = "BCoP"
)

func main() {
	client := &http.Client{
		Transport: wrapTransport(http.DefaultTransport.(*http.Transport)),
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, BCoPKey, "testtest\r\n")
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func wrapTransport(t *http.Transport) *http.Transport {
	t.DialContext = wrapDialContext(t.DialContext)
	return t
}

func wrapDialContext(dc func(ctx context.Context, network, addr string) (net.Conn, error)) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dc(ctx, network, addr)
		if err != nil {
			return nil, err
		}
		bag := ctx.Value(BCoPKey)
		// wrap conn
		conn.Write([]byte(bag.(string)))
		return conn, err
	}
}
