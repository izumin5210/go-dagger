package main

type CoffeeMakerProvider interface {
	Get() *CoffeeMaker
}
