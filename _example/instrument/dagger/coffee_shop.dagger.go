package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type DaggerCoffeeShop struct {
	provideHeaterProvider HeaterProvider
	thermosiphonProvider  ThermosiphonProvider
	providePumpProvider   PumpProvider
	coffeeMakerProvider   CoffeeMakerProvider
}

func buildCoffeeShop(builder *DaggerCoffeeShop_Builder) *DaggerCoffeeShop {
	coffeeShop := &DaggerCoffeeShop{}
	coffeeShop.provideHeaterProvider = &Heater_DoubleCheckProvider{
		provider: &DripCoffeeModule_ProvideHeaterFactory{
			module: builder.dripCoffeeModule,
		},
	}
	coffeeShop.thermosiphonProvider = &Thermosiphon_Factory{
		heaterProvider: coffeeShop.provideHeaterProvider,
	}
	coffeeShop.providePumpProvider = &Thermosiphon_Pump_Factory{
		thermosiphonProvider: coffeeShop.thermosiphonProvider,
	}
	coffeeShop.coffeeMakerProvider = &CoffeeMaker_Factory{
		heaterProvider: coffeeShop.provideHeaterProvider,
		pumpProvider:   coffeeShop.providePumpProvider,
	}
	return coffeeShop
}

func (c *DaggerCoffeeShop) Maker() *instrument.CoffeeMaker {
	return &instrument.CoffeeMaker{
		Heater: c.provideHeaterProvider.Get(),
		Pump:   c.providePumpProvider.Get(),
	}
}

type DaggerCoffeeShop_Builder struct {
	dripCoffeeModule *instrument.DripCoffeeModule
}

func (b *DaggerCoffeeShop_Builder) Build() instrument.CoffeeShop {
	if b.dripCoffeeModule == nil {
		b.dripCoffeeModule = &instrument.DripCoffeeModule{}
	}
	return buildCoffeeShop(b)
}

func (b *DaggerCoffeeShop_Builder) DripCoffeeModule(dripCoffeeModule *instrument.DripCoffeeModule) *DaggerCoffeeShop_Builder {
	b.dripCoffeeModule = dripCoffeeModule
	return b
}
