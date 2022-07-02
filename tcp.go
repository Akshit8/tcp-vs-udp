package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
)

const EXIT_CMD = "exit"

func read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer

	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		buffer.Write(data)

		if !isPrefix {
			break
		}
	}

	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	n, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}

	return n, err
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	log.Printf("new connection from %v\n", conn.RemoteAddr())

	for {
		content, err := read(conn)
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}

		if content == EXIT_CMD {
			log.Printf("received exit command, exiting\n")
			return
		}

		log.Printf("received: %v\n", content)

		n, err := write(conn, content)
		if err != nil {
			log.Printf("write error: %v\n", err)
			return
		}

		log.Printf("sent %d bytes to conn: %v\n", n, conn.RemoteAddr())
	}
}

func tcp() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// using log.Fatal will exit the server
			log.Printf("accept error: %v", err)
		}

		go handleConn(conn)
	}
}
