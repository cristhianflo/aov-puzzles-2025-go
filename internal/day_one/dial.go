package day_one

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
