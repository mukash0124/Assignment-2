package main

import "fmt"

func main() {
	teslaS := getCar("Tesla Model S")
	teslaX := getCar("Tesla Model X")

	fmt.Printf("Price of %s is %d\n", teslaS.getModel(), teslaS.getPrice())
	fmt.Printf("Price of %s is %d\n", teslaX.getModel(), teslaX.getPrice())
}