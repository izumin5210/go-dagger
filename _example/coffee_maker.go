package main

import "fmt"

type CoffeeMaker struct {
	Heater Heater `dagger:"inject"`
	Pump   Pump   `dagger:"inject"`
}

func (m *CoffeeMaker) Brew() {
	m.Heater.On()
	m.Pump.Pump()
	fmt.Println(" [_]P coffee! [_]P ")
	m.Heater.Off()
}
