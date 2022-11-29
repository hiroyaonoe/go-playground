package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
