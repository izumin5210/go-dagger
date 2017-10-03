package main

type Heater interface {
	On()
	Off()
	IsHot() bool
}
