package main

type Thermosiphon_Factory struct {
	heaterProvider HeaterProvider
}

func (f *Thermosiphon_Factory) Get() *Thermosiphon {
	return &Thermosiphon{
		Heater: f.heaterProvider.Get(),
	}
}
