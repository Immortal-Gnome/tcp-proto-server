package main

import (
	"encoding/binary"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"tcp-proto-server/grid"
)

const PORT = ":7000"

var _grid *grid.Grid

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
	// Get client address information
	clientAddr := conn.RemoteAddr().String()

	// Log new connection
	log.Printf("Info: new client connected from %s\n", clientAddr)

	// Handle client connection
	buffer := make([]byte, 4) // Buffer size for an int32
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("Info: client %s disconnected\n", clientAddr)
			} else {
				log.Printf("ERROR: reading from client %s (%v)\n", clientAddr, err)
			}
			break
		}

		// Process the received int
		if n == 4 { // Expecting 4 bytes for an int32
			// Convert the 4 bytes to an integer
			receivedInt := int32(binary.LittleEndian.Uint32(buffer[:4]))
			//log.Printf("Received integer %d from client %s\n", receivedInt, clientAddr)
			if receivedInt == 1 {
				log.Println("Creating Grid")
				//_create_grid()
			}
			if receivedInt == 2 {
				log.Println("Coloring Tile")
			}
			if receivedInt == 3 {
				log.Println("Coloring All Tiles")
			}
			if receivedInt == 4 {
				log.Println("Clear All Tiles")
			}
			//log.Printf("Raw bytes received: %v", buffer[:n])
		} else {
			log.Printf("Warning: Received unexpected data size: %d bytes from %s\n", n, clientAddr)
		}
	}

	// Close the connection when done
	conn.Close()
}

func _create_grid() {
	if _grid == nil {
		log.Println("grid already defined")
		return
	}
	_grid = grid.New(5, 5)
}

func _color_random_tile() {
	x := rand.Intn(_grid.Width())
	y := rand.Intn(_grid.Height())
	var _ grid.Tile
	_, _ = _grid.Get(x, y)

	/* 	// Use the tile by setting its color
	   	color := grid.NewColor()
	   	color = grid.SetRandom(color)

	   	// Update the tile in the grid with the new color
	   	_grid.SetColor(x, y, *color) */

}

// Old Handle Connections Message, keeping this here for now

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

/* func handleConnection(conn net.Conn) {
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
*/
