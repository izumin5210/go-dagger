package main

import (
	"sync"
)

type Heater_DoubleCheckProvider struct {
	instance    Heater
	provider    HeaterProvider
	initialized bool
	sync.Mutex
}

func (p *Heater_DoubleCheckProvider) Get() Heater {
	if !p.initialized {
		p.Lock()
		defer p.Unlock()

		p.instance = p.provider.Get()
		p.provider = nil
		p.initialized = true
	}
	return p.instance
}
