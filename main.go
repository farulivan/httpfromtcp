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

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		currentLine := []byte{}
		for {
			slice := make([]byte, 8)
			n, err := f.Read(slice)
			if n > 0 {
				for _, b := range slice[:n] {
					if b == '\n' {
						ch <- string(currentLine)
						currentLine = currentLine[:0]
					} else {
						currentLine = append(currentLine, b)
					}
				}
			}
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
				break
			}
		}
		if len(currentLine) > 0 {
			ch <- string(currentLine)
		}
	}()
	return ch
}
