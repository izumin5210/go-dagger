package main

import "fmt"

type Thermosiphon struct {
	Heater Heater `dagger:"inject"`
}

func (t *Thermosiphon) Pump() {
	if t.Heater.IsHot() {
		fmt.Println("=> => pumping => =>")
	}
}
