package examplemetric

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat-community/module/examplemodule"
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	Metric.Register()

}

// Metric Setup
var Metric = helper.NewMetric("examplemetric", ExampleMetric{}, examplemodule.Module)

// Metricer object
type ExampleMetric struct {
	helper.MetricConfig
}

func (b ExampleMetric) Setup() {
}

// Fetch expvars from a running beat
func (b ExampleMetric) Fetch() (events []common.MapStr) {

	event := common.MapStr{
		"example": "hello world",
	}

	events = append(events, event)

	return events
}
