package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"tcp-proto-server/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("new connection from %s", conn.RemoteAddr())
	defer conn.Close()

	msg := data.Data{Value: 1, Timestamp: 0}
	d, err := proto.Marshal(&msg)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return
	}

	length, err := conn.Write(d)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return
	}

	log.Printf("wrote %d bytes", length)
}
