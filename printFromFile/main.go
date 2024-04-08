package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file_content, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file_content)
}
