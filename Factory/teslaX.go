package main

type TeslaX struct {
	Car
}

func newTeslaX() ICar {
	return &TeslaX{
		Car: Car {
			model: "Tesla Model X",
			price: 100000,
		},
	}
}