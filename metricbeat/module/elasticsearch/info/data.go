package info

import (
	"github.com/elastic/beats/libbeat/common"
)

type Data struct {
	Name        string  `json:"name"`
	ClusterName string  `json:"cluster_name"`
	Version     Version `json:"version"`
}

type Version struct {
	Number         string `json:"number"`
	BuildHash      string `json:"build_hash"`
	BuildTimestamp string `json:"build_timestamp"`
	BuildSnapshot  bool   `json:"build_snapshot"`
	LuceneVersion  string `json:"lucene_version"`
}

// Map data to MapStr
func (data *Data) eventMapping() common.MapStr {

	// TODO: Is there a way to automate this if it maps 1-1 the data object?
	event := common.MapStr{
		"name":         data.Name,
		"cluster_name": data.ClusterName,
		"version": common.MapStr{
			"number":          data.Version.Number,
			"build_hash":      data.Version.BuildHash,
			"build_timestamp": data.Version.BuildTimestamp,
			"build_snapshot":  data.Version.BuildSnapshot,
			"lucene_version":  data.Version.LuceneVersion,
		},
	}

	return event
}
