package main

import (
	"github.com/izumin5210/go-dagger/_example/instrument"
	"github.com/izumin5210/go-dagger/_example/instrument/dagger"
)

func main() {
	shop := (&dagger.DaggerCoffeeShop_Builder{}).
		DripCoffeeModule(&instrument.DripCoffeeModule{}).
		Build()
	shop.Maker().Brew()
}
