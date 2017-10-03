package instrument

//dagger:module
type PumpModule interface {
	ProvidePump(pump Thermosiphon) Pump
}
