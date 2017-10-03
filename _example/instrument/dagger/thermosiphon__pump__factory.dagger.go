package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type Thermosiphon_Pump_Factory struct {
	thermosiphonProvider ThermosiphonProvider
}

func (f *Thermosiphon_Pump_Factory) Get() instrument.Pump {
	return f.thermosiphonProvider.Get()
}
