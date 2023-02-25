package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("starting Protohackers p000")

	ln, err := net.Listen("tcp", ":18001")
	if err != nil {
		panic(err)
	}

	log.Println("listening on port 18001")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept: %v", err)
			continue
		}

		conn.SetDeadline(time.Now().Add(5 * time.Second))

		log.Printf("got conn from: %s", conn.RemoteAddr())

		go func() {
			if err := handle(conn); err != nil {
				log.Printf("handle: %v", err)
			}
		}()
	}
}

func handle(conn net.Conn) error {
	defer conn.Close()

	n, err := io.Copy(conn, conn)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	log.Printf("for %s written %d", conn.RemoteAddr(), n)

	return nil
}
