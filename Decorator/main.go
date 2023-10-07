package main

import "fmt"

func main() {
	tesla := &Tesla{}
	ferrari := &Ferrari{}

	teslaWithParktronics := &Parktronic{
		car: tesla,
	}

	ferrariWithClimateControl := &ClimateControl{
		car: ferrari,
	}

	fmt.Printf("Price of car of model Tesla with parktronics is %d\n", teslaWithParktronics.getPrice())
	fmt.Printf("Price of car of model Ferrari with climate-control is %d\n", ferrariWithClimateControl.getPrice())
}