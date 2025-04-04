package grid

import (
	"math/rand"
)

type Color struct {
	r float32
	g float32
	b float32
}

func NewColor() *Color {
	return &Color{
		r: 0.0,
		g: 0.0,
		b: 0.0,
	}
}

func set_random(c *Color) *Color {
	c.r = rand.Float32()
	c.g = rand.Float32()
	c.b = rand.Float32()
	return c
}

func set_white(c *Color) *Color {
	c.r = 1.0
	c.g = 1.0
	c.b = 1.0
	return c
}
