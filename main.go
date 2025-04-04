package main

import (
	"log"
	"net"
	"os"
	data "tcp-proto-server/proto"
	"time"

	"google.golang.org/protobuf/proto"
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

/* func handleConnection(conn net.Conn) {
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
} */

func handleConnection(conn net.Conn) {
	log.Printf("Info: new connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	// Create a buffer to read incoming data
	buffer := make([]byte, 1024)

	// Read data from the connection
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("ERROR: could not read from connection (%v)\n", err)
		return
	}

	log.Printf("Info: read %d bytes\n", n)

	// Unmarshal the received protobuf message
	var receivedMsg data.Data
	if err := proto.Unmarshal(buffer[:n], &receivedMsg); err != nil {
		log.Printf("ERROR: could not decode message (%v)\n", err)
		return
	}

	// Process the received message
	log.Printf("Info: received message with value=%d and timestamp=%d\n",
		receivedMsg.Value, receivedMsg.Timestamp)

	// Example: Update the value and send response back
	responseMsg := data.Data{
		Value:     receivedMsg.Value * 2, // Example processing: double the value
		Timestamp: time.Now().Unix(),
	}

	// Marshal the response message
	responseData, err := proto.Marshal(&responseMsg)
	if err != nil {
		log.Printf("ERROR: could not encode response message (%v)\n", err)
		return
	}

	// Send the response back to the client
	length, err := conn.Write(responseData)
	if err != nil {
		log.Printf("ERROR: could not write response to connection (%v)\n", err)
		return
	}

	log.Printf("Info: wrote %d bytes as response\n", length)
}
