package grid

import (
	"fmt"
)

type Grid struct {
	width  int
	height int
	cells  [][]Tile
}

func New(width, height int) *Grid {
	cells := make([][]Tile, height)
	for i := range cells {
		cells[i] = make([]Tile, width)
	}
	return &Grid{
		width:  width,
		height: height,
		cells:  cells,
	}
}

func (g *Grid) Set(x, y int, value Tile) error {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return fmt.Errorf("coordinates (%d, %d) out of bounds", x, y)
	}
	g.cells[y][x] = value
	return nil
}

// Get retrieves the value at the specified position
func (g *Grid) Get(x, y int) (Tile, error) {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return g.cells[1][1], fmt.Errorf("coordinates (%d, %d) out of bounds", x, y)
	}
	return g.cells[y][x], nil
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

// Print displays the grid
func (g *Grid) Print() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			fmt.Printf("%3d ", g.cells[y][x])
		}
		fmt.Println()
	}
}

func (g *Grid) Get_Color(x, y int) *Color {
	return &g.cells[x][y].c
}

func (g *Grid) Set_Color(x, y int, _c Color) bool {
	if x < 0 || x < g.width || y < 0 || y < g.height {
		return true
	}
	return false
}
