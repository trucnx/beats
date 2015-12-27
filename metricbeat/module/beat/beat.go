// Beat module and metric
package beat

import (
	"github.com/elastic/beats/metricbeat/helper"
)

// This one comabines module and metric
func init() {
	Module.Register()
}

// Module object
var Module = helper.NewModule("beat", Beat{})

type Beat struct {
}

func (b Beat) Setup() {

}
