package dagger

import "github.com/izumin5210/go-dagger/_example/instrument"

type CoffeeMakerProvider interface {
	Get() *instrument.CoffeeMaker
}
