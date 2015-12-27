// Fetches stats from http://elasticsearch-host/_stats
package stats

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

var Config TestMetricConfig

// Metric object
var Metric = helper.NewMetric("stats", Stats{}, elasticsearch.Module)

type Stats struct {
	helper.MetricConfig
	Config TestMetricConfig
}

func init() {
	Metric.Register()
}

type TestMetricConfig struct {
	Period string
}

func (ts Stats) Setup() {
}

func (ts Stats) Fetch() (events []common.MapStr) {

	for _, host := range elasticsearch.Config.Hosts {
		data := &Data{}
		elasticsearch.LoadUrl(host+"/_stats", data)

		events = append(events, data.eventMapping())
	}
	return events
}
