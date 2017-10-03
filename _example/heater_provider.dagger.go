package main

type HeaterProvider interface {
	Get() Heater
}
