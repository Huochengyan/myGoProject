package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var url = "https://h5.pipix.com/s/6mtP8N"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

}
