package main

func main() {
	shop := DaggerCoffeeShop.Builder{}.
		DripCoffeeModule(&DripCoffeeModule{}).
		Build()
	shop.Maker().Brew()
}
