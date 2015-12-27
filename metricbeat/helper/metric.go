package helper

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"gopkg.in/yaml.v2"
	"time"
)

// Base metric configuration
type MetricConfig struct {
	Period string
}

// Metric specific data
// This must be defined by each metric
type Metric struct {
	Name string
	// Metric specific config

	// Generic Config existing in all metrics
	BaseConfig MetricConfig

	// Raw metric specific config
	// This is provided to convert it into Config later
	RawConfig interface{}

	// Metric specific config
	Config interface{}

	Metricer Metricer
	Module   *Module
}

// Interface for each metric
type Metricer interface {
	// TODO: Setup can be optional, so we could make it a second interface we check for?
	Setup()
	Fetch() []common.MapStr
	// TODO: add errors to interface messages
	// TODO: Add stop method
}

func NewMetric(name string, metricer Metricer, module *Module) *Metric {
	return &Metric{
		Name:     name,
		Metricer: metricer,
		Module:   module,
	}
}

func (m *Metric) LoadConfig(config interface{}) {

	bytes, err := yaml.Marshal(m.RawConfig)

	if err != nil {
		logp.Err("Load metric config error: %v", err)
	}
	yaml.Unmarshal(bytes, config)
}

// Registers metric with module
func (m *Metric) Register() {
	m.Module.AddMetric(m)
}

// RunMetric runs the given metric
func (m *Metric) Run(b *beat.Beat) {

	m.Metricer.Setup()
	period, err := time.ParseDuration(m.BaseConfig.Period)

	if err != nil {
		logp.Info("Error in parsing period of metric %s: %v", m.Name, err)
	}

	// If no period set, set default
	if period == 0 {
		logp.Info("Setting default period for metric %s as not set.", m.Name)
		period = 1 * time.Second
	}

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	logp.Info("Start metric %s with period %v", m.Name, period)

	for {
		select {
		case <-ticker.C:
		}

		events := m.Metricer.Fetch()
		newEvents := []common.MapStr{}

		// Default names based on module and metric
		// These can be overwritten by setting index or / and type in the event
		indexName := m.Module.Name
		typeName := m.Name
		timestamp := common.Time(time.Now())

		for _, event := range events {
			// Set index from event if set
			if _, ok := event["index"]; ok {
				indexName = event["index"].(string)
				delete(event, "index")
			}

			// Set type from event if set
			if _, ok := event["type"]; ok {
				typeName = event["type"].(string)
				delete(event, "type")
			}

			// Set timestamp from event if set, move it to the top level
			// If not set, timestamp is created
			if _, ok := event["@timestamp"]; ok {
				timestamp = event["@timestamp"].(common.Time)
				delete(event, "@timestamp")
			}

			// TODO: Add root level option?
			event = common.MapStr{
				"index":      indexName,
				"type":       typeName,
				"metric":     event,
				"@timestamp": timestamp,
			}

			newEvents = append(newEvents, event)
		}

		b.Events.PublishEvents(newEvents)
	}
}
