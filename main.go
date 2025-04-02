package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"tcp-proto-server/proto"
)

const PORT = ":7000"

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Printf("ERROR: could not create tcp socket (%v)\n", err)
		os.Exit(1)
	}

	log.Printf("Info: listening on port %s\n", PORT)

	for {
		if conn, err := listener.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("Info: new connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	msg := data.Data{Value: 1, Timestamp: 0}
	d, err := proto.Marshal(&msg)
	if err != nil {
		log.Printf("ERROR: could not encode message (%v)\n", err)
		return
	}

	length, err := conn.Write(d)
	if err != nil {
		log.Printf("ERROR: could not write to connection (%v)\n", err)
		return
	}

	log.Printf("Info: wrote %d bytes\n", length)
}
