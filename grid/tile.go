package grid

type Tile struct {
	c *Color
}

func (t *Tile) Set(color *Color) {
	t.c = color
}
