package main

//dagger:module
type PumpModule interface {
	ProvidePump(pump Thermosiphon) Pump
}
