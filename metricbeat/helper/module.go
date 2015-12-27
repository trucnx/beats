package helper

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"gopkg.in/yaml.v2"
)

// Module specifics
// This must be defined by each module
type Module struct {
	Name    string
	Moduler Moduler

	Metrics map[string]*Metric
	// Generic config existing in all modules
	BaseConfig ModuleConfig

	// Raw module specific config
	// This is provided to convert it into Config later
	RawConfig interface{}

	// Module specific Config
	// TODO: Does not work yet because of local type
	Config interface{}
}

func NewModule(name string, moduler Moduler) *Module {
	return &Module{
		Name:    name,
		Moduler: moduler,
		Metrics: map[string]*Metric{},
	}
}

func (m *Module) Register() {
	Registry.AddModule(m)
}

// Add metric to module
func (m *Module) AddMetric(metric *Metric) {
	m.Metrics[metric.Name] = metric
}

// Interface for each module
type Moduler interface {
	// TODO: Setup can be optional, so we could make it a second interface we check for?
	Setup()
}

// Base configuration for list of modules
type ModulesConfig struct {
	Modules map[string]ModuleConfig
}

// Base module configuration
type ModuleConfig struct {
	Metrics map[string]MetricConfig
}

// Loads the configurations specific config.
// This needs the configuration object defined inside the module
func (m *Module) LoadConfig(config interface{}) {
	bytes, err := yaml.Marshal(m.RawConfig)

	if err != nil {
		logp.Err("Load module config error: %v", err)
	}
	yaml.Unmarshal(bytes, config)
}

// Starts the given module
func (module *Module) Start(b *beat.Beat) {

	logp.Info("Start Module: %v", module.Name)

	module.Moduler.Setup()

	for _, metric := range module.Metrics {
		// TODO: If a metric panics, it should not affect other modules
		go metric.Run(b)
	}
}
