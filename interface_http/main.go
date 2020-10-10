package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lg := logger{}

	io.Copy(lg, resp.Body)
}

func (logger) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("Just logged this many bytes", len(p))

	return len(p), nil
}
