package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const serverAddr = "localhost:42069"

func main() {
	// Resolve the UDP address
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatalf("Error resolving UDP address: %s\n", err)
	}

	// Dial the UDP address
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error dialing UDP address: %s\n", err)
	}
	defer conn.Close()

	// Create a reader for stdin.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			fmt.Println("\nExiting...")
			return
		}
		if err != nil {
			log.Printf("Error reading from stdin: %s\n", err)
			continue
		}
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to UDP connection: %s\n", err)
			continue
		}

		fmt.Println("Message sent: ", message)
	}
}
