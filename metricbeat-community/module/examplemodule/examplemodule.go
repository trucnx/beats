package examplemodule

import (
	"github.com/elastic/beats/metricbeat/helper"
)

// This one combines module and metric
func init() {
	Module.Register()
}

// Module object
var Module = helper.NewModule("examplemodule", Beat{})

type Beat struct {
}

func (b Beat) Setup() {

}
