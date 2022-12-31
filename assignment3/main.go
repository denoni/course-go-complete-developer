package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(os.Stdout, f)
}
