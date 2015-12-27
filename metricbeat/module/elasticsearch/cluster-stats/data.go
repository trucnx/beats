package cluster_stats

import (
	"github.com/elastic/beats/libbeat/common"
)

// Data mapping of retrieved data
type Data struct {
	ClusterName string  `json:"cluster_name"`
	Status      string  `json:"status"`
	Indicies    Indices `json:"indices"`
	Nodes       Nodes   `json:"nodes"`
}

type Indices struct {
	Count  int    `json:"count"`
	Shards Shards `json:"shards"`
}

type Shards struct {
	Total     int `json:"total"`
	Primaries int `json:"primaries"`
}

type Nodes struct {
	Count Count
}

type Count struct {
	Total      int `json:"total"`
	MasterOnly int `json:"master_only"`
	DataOnly   int `json:"data_only"`
	MasterData int `json:"master_data"`
	Client     int `json:"client"`
}

// Map data to MapStr
func (data *Data) eventMapping() common.MapStr {

	// TODO: Is there a way to automate this if it maps 1-1 the data object?
	event := common.MapStr{
		"name":   data.ClusterName,
		"status": data.Status,
		"indices": common.MapStr{
			"count": data.Indicies.Count,
			"shards": common.MapStr{
				"total":     data.Indicies.Shards.Total,
				"primaries": data.Indicies.Shards.Primaries,
			},
		},
		"nodes": common.MapStr{
			"count": common.MapStr{
				"total":       data.Nodes.Count.Total,
				"master_only": data.Nodes.Count.MasterOnly,
				"data_only":   data.Nodes.Count.DataOnly,
				"master_data": data.Nodes.Count.MasterData,
				"client":      data.Nodes.Count.Client,
			},
		},
	}

	return event
}
