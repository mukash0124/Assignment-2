package main

type Parktronic struct {
	car ICar
}

func (c *Parktronic) getPrice() int {
	return c.car.getPrice() + 5000
}