package main

type Car struct {
	model string
	price int
}

func (c *Car) getModel() string {
	return c.model
}
func (c *Car) getPrice() int {
	return c.price
}