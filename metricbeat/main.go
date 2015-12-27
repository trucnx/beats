package main

import (
	metricbeat "github.com/elastic/beats/metricbeat/beat"

	"github.com/elastic/beats/libbeat/beat"
	//"github.com/elastic/beats/metricbeat/module"
	//"github.com/elastic/beats/metricbeat/module/all"
)

// You can overwrite these, e.g.: go build -ldflags "-X main.Version 1.0.0-beta3"
var Version = "1.0.0"
var Name = "metricbeat"

func main() {
	beat.Run(Name, Version, metricbeat.New())

}
