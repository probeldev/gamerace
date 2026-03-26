package model

type Car struct {
	X       int
	Y       int
	VisualY float64
}

func (c *Car) Down(speed int) {
	c.VisualY += 1 / float64(speed)

	if float64(c.Y) < c.VisualY {
		c.Y = int(c.VisualY)
	}
}
