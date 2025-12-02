package day_one

import (
	"math"
)

type Dial struct {
	value int
	size  int
}

func NewDial(size int) *Dial {
	return &Dial{
		value: 0,
		size:  size,
	}
}

func (c *Dial) Value() int {
	return c.value
}

func (c *Dial) Set(v int) {
	c.value = ((v % c.size) + c.size) % c.size
}

func (c *Dial) Rotate(delta int) {
	c.value = ((c.value+delta)%c.size + c.size) % c.size
}

func (c *Dial) CalculateRotations(v int, direction string) int {
	base := c.value

	if direction == "L" && c.value != 0 {
		base = c.size - c.value
	}
	sum := base + v

	return int(math.Abs(float64(sum)) / float64(c.size))
}
