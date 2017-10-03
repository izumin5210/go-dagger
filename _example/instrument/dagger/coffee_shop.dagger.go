package main

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

func (c *DaggerCoffeeShop) Maker() *CoffeeMaker {
	return &CoffeeMaker{
		Heater: c.provideHeaterProvider.Get(),
		Pump:   c.providePumpProvider.Get(),
	}
}

type DaggerCoffeeShop_Builder struct {
	dripCoffeeModule *DripCoffeeModule
}

func (b *DaggerCoffeeShop_Builder) Build() CoffeeShop {
	if b.dripCoffeeModule == nil {
		b.dripCoffeeModule = &DripCoffeeModule{}
	}
	return buildCoffeeShop(b)
}

func (b *DaggerCoffeeShop_Builder) DripCoffeeModule(dripCoffeeModule *DripCoffeeModule) *DaggerCoffeeShop_Builder {
	b.dripCoffeeModule = dripCoffeeModule
	return b
}
