package dagger

import (
	"sync"

	"github.com/izumin5210/go-dagger/_example/instrument"
)

type Heater_DoubleCheckProvider struct {
	instance    instrument.Heater
	provider    HeaterProvider
	initialized bool
	sync.Mutex
}

func (p *Heater_DoubleCheckProvider) Get() instrument.Heater {
	if !p.initialized {
		p.Lock()
		defer p.Unlock()

		p.instance = p.provider.Get()
		p.provider = nil
		p.initialized = true
	}
	return p.instance
}
