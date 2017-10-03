package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type DripCoffeeModule_ProvideHeaterFactory struct {
	module *instrument.DripCoffeeModule
}

func (f *DripCoffeeModule_ProvideHeaterFactory) Get() instrument.Heater {
	return f.module.ProvideHeater()
}
