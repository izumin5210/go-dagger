package main

import (
	"fmt"
)

type ElectricHeater struct {
	heating bool
}

func (h *ElectricHeater) On() {
	fmt.Println("~ ~ ~ heating ~ ~ ~")
	h.heating = true
}

func (h *ElectricHeater) Off() {
	h.heating = false
}

func (h *ElectricHeater) IsHot() bool {
	return h.heating
}
