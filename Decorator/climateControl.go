package main

type ClimateControl struct {
	car ICar
}

func (c *ClimateControl) getPrice() int {
	return c.car.getPrice() + 4000;
}