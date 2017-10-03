package main

type Thermosiphon_Pump_Factory struct {
	thermosiphonProvider ThermosiphonProvider
}

func (f *Thermosiphon_Pump_Factory) Get() Pump {
	return f.thermosiphonProvider.Get()
}
