package elasticsearch

import (
	"encoding/json"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"io/ioutil"
	"net/http"
)

func init() {
	Module.Register()
}

var Module = helper.NewModule("elasticsearch", Elasticsearch{})

var Config = &ElasticsearchModuleConfig{}

type ElasticsearchModuleConfig struct {
	Metrics map[string]interface{}
	Hosts   []string
}

type Elasticsearch struct {
	Name    string
	Config  ElasticsearchModuleConfig
	Metrics map[string]helper.Metricer
}

func (e Elasticsearch) Setup() {

	// Loads module config
	// This is module specific config object
	Module.LoadConfig(&Config)
}

func LoadUrl(url string, data interface{}) {

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, data)
	if err != nil {
		logp.Err("Json error: %v", err)
	}
}
