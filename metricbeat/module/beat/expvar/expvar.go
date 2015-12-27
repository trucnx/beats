// Fetches statistics from the output of running beats
package expvar

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/beat"
)

func init() {
	Metric.Register()

}

// Metric Setup
var Metric = helper.NewMetric("expvar", BeatMetric{}, beat.Module)

// Metricer object
type BeatMetric struct {
	helper.MetricConfig
}

func (b BeatMetric) Setup() {
}

// Fetch expvars from a running beat
func (b BeatMetric) Fetch() (events []common.MapStr) {

	//path := "http://localhost:6060/debug/vars"

	event := common.MapStr{
		"type":  "helloworld",
		"index": "indexnameyes",
		"bac":   "rrre",
	}

	events = append(events, event)

	return events
}
