package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Error opening file %s: %s\n", inputFilePath, err)
	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	for {
		slice := make([]byte, 8)
		n, err := f.Read(slice)
		if n > 0 {
			fmt.Printf("read: %s\n", string(slice[:n]))
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
			break
		}
	}
}
