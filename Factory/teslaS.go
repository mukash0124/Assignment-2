package main

type TeslaS struct {
	Car
}

func newTeslaS() ICar {
	return &TeslaS{
		Car: Car {
			model: "Tesla Model S",
			price: 90000,
		},
	}
}