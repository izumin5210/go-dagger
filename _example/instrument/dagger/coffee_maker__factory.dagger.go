package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type CoffeeMaker_Factory struct {
	heaterProvider HeaterProvider
	pumpProvider   PumpProvider
}

func (f *CoffeeMaker_Factory) Get() *instrument.CoffeeMaker {
	return &instrument.CoffeeMaker{
		Heater: f.heaterProvider.Get(),
		Pump:   f.pumpProvider.Get(),
	}
}
