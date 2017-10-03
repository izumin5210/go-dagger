package main

type CoffeeMaker_Factory struct {
	heaterProvider HeaterProvider
	pumpProvider   PumpProvider
}

func (f *CoffeeMaker_Factory) Get() *CoffeeMaker {
	return &CoffeeMaker{
		Heater: f.heaterProvider.Get(),
		Pump:   f.pumpProvider.Get(),
	}
}
