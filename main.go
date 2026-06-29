package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listen for TCP traffic on port %s: %s\n", port, err)
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s\n", err)
			continue
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		fmt.Println("=====================================")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Println(line)
		}
		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
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
