// Fetches basic version info about elasticsearch hosts from http://elasticsearchhost/
package info

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

var Config TestMetricConfig

// Metric object
var Metric = helper.NewMetric("info", Info{}, elasticsearch.Module)

type Info struct {
	helper.MetricConfig
	Config TestMetricConfig
}

func init() {
	Metric.Register()
}

type TestMetricConfig struct {
	Period string
}

func (info Info) Setup() {
}

func (info Info) Fetch() (events []common.MapStr) {

	for _, host := range elasticsearch.Config.Hosts {
		data := &Data{}
		elasticsearch.LoadUrl(host, data)

		events = append(events, data.eventMapping())
	}
	return events
}
