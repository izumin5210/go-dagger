package main

//dagger:modulle
type PumpModule interface {
	ProvidePump(pump Thermosiphon) Pump
}
