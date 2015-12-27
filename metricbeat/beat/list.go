package beat

// Make sure all active plugins are loaded
// TODO: create a script to automatically generate this list
import (
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"

	// List of all metrics to make sure they are registred
	// Every new metric must be added here
	_ "github.com/elastic/beats/metricbeat/module/beat"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/cluster-stats"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/info"
	_ "github.com/elastic/beats/metricbeat/module/elasticsearch/stats"
	_ "github.com/elastic/beats/metricbeat/module/redis/info"
)

func ListAll() {
	logp.Debug("beat", "Registered Modules and Metrics")
	for moduleName, module := range helper.Registry {
		for metricName, _ := range module.Metrics {
			logp.Debug("beat", "Registred: Module: %v, Metric: %v", moduleName, metricName)
		}
	}
}
