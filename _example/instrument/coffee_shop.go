package instrument

//dagger:singleton
//dagger:component(modules = {DripCoffeeModule})
type CoffeeShop interface {
	Maker() *CoffeeMaker
}
