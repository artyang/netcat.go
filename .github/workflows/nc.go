package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]
	target := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %s\n", target, err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", target)

	go readConnection(conn)
	writeConnection(conn)
}

func readConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading data: %s\n", err)
			return
		}

		fmt.Print(message)
	}
}

func writeConnection(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err)
			return
		}

		_, err = writer.WriteString(message)
		if err != nil {
			fmt.Printf("Error writing data: %s\n", err)
			return
		}

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing data: %s\n", err)
			return
		}
	}
}
