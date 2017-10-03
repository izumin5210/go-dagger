package main

type PumpProvider interface {
	Get() Pump
}
