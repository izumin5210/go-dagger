package main

//dagger:modulle(includes = PumpModule.class)
type DripCoffeeModule struct {
}

//dagger:provide
func (m *DripCoffeeModule) ProvideHeater() Heater {
	return &ElectricHeater{}
}
