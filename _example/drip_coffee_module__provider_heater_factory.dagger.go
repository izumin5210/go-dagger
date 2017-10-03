package main

type DripCoffeeModule_ProvideHeaterFactory struct {
	module *DripCoffeeModule
}

func (f *DripCoffeeModule_ProvideHeaterFactory) Get() Heater {
	return f.module.ProvideHeater()
}
