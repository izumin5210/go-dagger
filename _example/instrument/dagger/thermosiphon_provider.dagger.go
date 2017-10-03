package dagger

import (
	"github.com/izumin5210/go-dagger/_example/instrument"
)

type ThermosiphonProvider interface {
	Get() *instrument.Thermosiphon
}
