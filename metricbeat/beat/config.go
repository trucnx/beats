package beat

import (
	"github.com/elastic/beats/metricbeat/helper"
)

type MetricbeatConfig struct {
	Metricbeat helper.ModulesConfig
}

// Raw module config to be processed later by the module
type RawModulesConfig struct {
	Metricbeat struct {
		Modules map[string]interface{}
	}
}

// Raw metric config to be processed later by the metric
type RawMetricsConfig struct {
	Metricbeat struct {
		Modules map[string]struct {
			Metrics map[string]interface{}
		}
	}
}
