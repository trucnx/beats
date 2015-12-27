package helper

import (
	"github.com/elastic/beats/libbeat/beat"
)

// Global register for modules and metrics
// Each module name must be unique
// Each module-metric name combination must be unique
// TODO: Should this moved into the metricbeat object? Not possible because of init?
var Registry = Register{}

type Register map[string]*Module

func StartModules(b *beat.Beat) {
	for _, module := range Registry {
		// TODO: If a module panics, it should not affect other modules
		go module.Start(b)
	}
}

func (r Register) AddModule(m *Module) {
	r[m.Name] = m
}
