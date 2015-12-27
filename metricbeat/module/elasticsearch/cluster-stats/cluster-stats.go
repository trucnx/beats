// Fetches the cluster statistics from the given elasticsearch hosts
//
// Calls url http://eshost/_cluster/stats
package cluster_stats

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/elasticsearch"
)

// Adds metric to module
func init() {
	Metric.Register()
}

// Metric object
var Metric = helper.NewMetric("cluster-stats", ClusterStats{}, elasticsearch.Module)

// Detailed configuration
var Config = ClusterStatsConfig{}

type ClusterStatsConfig struct {
	Test string
}

type ClusterStats struct {
	helper.MetricConfig
}

func (cs ClusterStats) Setup() {
	// Loads module config
	// This is module specific config object
	Metric.LoadConfig(&Config)
}

// TODO Should we use a pointer here?
func (cs ClusterStats) Fetch() (events []common.MapStr) {

	for _, host := range elasticsearch.Config.Hosts {
		data := &Data{}
		elasticsearch.LoadUrl(host+"/_cluster/stats", data)

		events = append(events, data.eventMapping())
	}
	return events
}
