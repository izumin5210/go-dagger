package instrument

//dagger:module
type DripCoffeeModule struct {
	PumpModule
}

//dagger:provide
func (m *DripCoffeeModule) ProvideHeater() Heater {
	return &ElectricHeater{}
}
