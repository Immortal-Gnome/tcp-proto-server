package main

import (
	"encoding/binary"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"tcp-proto-server/grid"
	data "tcp-proto-server/proto"

	"google.golang.org/protobuf/proto"
)

const PORT = ":7000"

var _grid *grid.Grid
var pending_grid_update bool

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

	var update_package *data.Grid_Data = nil
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

		if n == 4 {
			receivedInt := int32(binary.LittleEndian.Uint32(buffer[:4]))
			if receivedInt == 1 {
				log.Println("\nCreating Grid\n")
				_create_grid()
				pending_grid_update = true
			}

			if receivedInt == 2 {
				log.Println("\nColoring Tile\n")
				_color_random_tile()
				pending_grid_update = true
			}
			if receivedInt == 3 {
				log.Println("\nColoring All Tiles\n")
				_color_all_tiles()
				pending_grid_update = true
			}

			if receivedInt == 4 {
				log.Println("Clear All Tiles")
				_clear_all_tiles()
				pending_grid_update = true
			}

			if receivedInt == 5 {
				_grid.Print()

			}

		} else {
			log.Printf("Warning: Received unexpected data size: %d bytes from %s\n", n, clientAddr)
		}

		if pending_grid_update {
			update_package = prepare_grid_update_package()
			transmit_grid_update(conn, update_package)
			pending_grid_update = false
			log.Println("Pending_Update_DONE")
		}
	}

	conn.Close()
}

func prepare_grid_update_package() *data.Grid_Data {

	_grid_data := new(data.Grid_Data)

	for x := 0; x < _grid.Width(); x++ {
		for y := 0; y < _grid.Height(); y++ {
			color := _grid.Get_Color(x, y)
			_tile_data := data.Tile_Data{
				R: float64(*color.Red()),
				G: float64(*color.Green()),
				B: float64(*color.Blue()),
				X: int32(x),
				Y: int32(y),
			}
			_grid_data.Tiles = append(_grid_data.Tiles, &_tile_data)
		}
	}
	log.Println("preparing big pile of shit")
	return _grid_data
}

func transmit_grid_update(conn net.Conn, data *data.Grid_Data) {
	// Marshal the response message
	responseData, err := proto.Marshal(data)
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

func _create_grid() {
	if _grid != nil {
		log.Println("grid already defined")
		return
	}
	_grid = grid.New(5, 5)
}

func _color_random_tile() {
	x := rand.Intn(_grid.Width() - 1)
	y := rand.Intn(_grid.Height() - 1)
	var tile grid.Tile
	tile, _ = _grid.Get(x, y)

	color := grid.NewColor()
	color = grid.SetRandom(color)
	tile.Set(color)
	_grid.Set(x, y, tile)

}

func _color_all_tiles() {
	for x := 0; x < _grid.Width(); x++ {
		for y := 0; y < _grid.Height(); y++ {
			var tile grid.Tile
			tile, _ = _grid.Get(x, y)
			color := grid.NewColor()
			color = grid.SetRandom(color)
			tile.Set(color)
			_grid.Set(x, y, tile)
		}
	}
}

func _clear_all_tiles() {
	for x := 0; x < _grid.Width(); x++ {
		for y := 0; y < _grid.Height(); y++ {
			var tile grid.Tile
			tile, _ = _grid.Get(x, y)
			color := grid.NewColor()
			color = grid.SetWhite(color)
			tile.Set(color)
			_grid.Set(x, y, tile)
		}
	}
}

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
