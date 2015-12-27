// Loads data from the redis info command
package info

import (
	//"fmt"
	"github.com/elastic/beats/libbeat/common"
	//"github.com/elastic/beats/libbeat/logp"
	rd "github.com/garyburd/redigo/redis"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/module/redis"
	"strings"
)

func init() {
	Metric.Register()
}

// Metric object
var Metric = helper.NewMetric("info", Redis{}, redis.Module)

var Config = &RedisMetricConfig{}

type RedisMetricConfig struct {
}

type Redis struct {
	Name   string
	Config RedisMetricConfig
}

func (e Redis) Setup() {
	// Loads module config
	// This is module specific config object
	Metric.LoadConfig(&Config)
}

func (e Redis) Fetch() (events []common.MapStr) {

	conn := redis.Connect()

	out, err := rd.String(conn.Do("INFO"))
	if err != nil {
		logp.Err("Error converting to string: %v", err)
	}

	// Feed every line into
	result := strings.Split(out, "\r\n")

	// Load redis info values into array
	values := map[string]string{}

	for _, value := range result {
		// Values are separated by :
		parts := strings.Split(value, ":")
		if len(parts) == 2 {
			values[parts[0]] = parts[1]
		}
	}

	event := common.MapStr{
		"version":    values["redis_version"],
		"mode":       values["redis_mode"],
		"os":         values["os"],
		"process_id": values["process_id"],
	}

	// All available values
	//logp.Debug("redis","Values: %+v", values)

	return append(events, event)
}
