package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs_body := make([]byte, 999999) //takes byte slice and initializes with 999999 empty spaces
	// resp.Body.Read(bs_body)
	// fmt.Println(string(bs_body))

	// io.Copy(os.Stdout, resp.Body) // (writer, reader)

	lw := logWriter{}
	io.Copy(lw, resp.Body) // (writer, reader)

}

// Custom Write method replicating the Write method in Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("\nJust wrote these many bytes:", len(bs))
	// check documentation for Writer to check what is supposed to be returned
	return len(bs), nil
}
