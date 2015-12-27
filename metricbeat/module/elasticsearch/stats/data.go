package stats

import (
	"github.com/elastic/beats/libbeat/common"
)

type Data struct {
	Shards Shards `json:"_shards"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

// Map data to MapStr
func (data *Data) eventMapping() common.MapStr {

	event := common.MapStr{
		"shards": common.MapStr{
			"total":      data.Shards.Total,
			"successful": data.Shards.Successful,
			"failed":     data.Shards.Failed,
		},
	}

	return event
}
