package main

func main() {
	shop := (&DaggerCoffeeShop_Builder{}).
		DripCoffeeModule(&DripCoffeeModule{}).
		Build()
	shop.Maker().Brew()
}
