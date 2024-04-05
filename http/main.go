package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs_body := make([]byte, 999999) //takes byte slice and initializes with 999999 empty spaces
	// resp.Body.Read(bs_body)
	// fmt.Println(string(bs_body))

	io.Copy(os.Stdout, resp.Body)
}
