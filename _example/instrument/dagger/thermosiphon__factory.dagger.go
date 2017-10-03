package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type Thermosiphon_Factory struct {
	heaterProvider HeaterProvider
}

func (f *Thermosiphon_Factory) Get() *instrument.Thermosiphon {
	return &instrument.Thermosiphon{
		Heater: f.heaterProvider.Get(),
	}
}
