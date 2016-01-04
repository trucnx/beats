package beat

import (
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
)

type Metricbeat struct {
	done          chan struct{}
	MbConfig      *MetricbeatConfig
	ModulesConfig *RawModulesConfig
	MetricsConfig *RawMetricsConfig
}

func New() *Metricbeat {
	return &Metricbeat{}
}

func (mb *Metricbeat) Config(b *beat.Beat) error {

	mb.MbConfig = &MetricbeatConfig{}
	err := cfgfile.Read(mb.MbConfig, "")
	if err != nil {
		fmt.Println(err)
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	mb.ModulesConfig = &RawModulesConfig{}
	err = cfgfile.Read(mb.ModulesConfig, "")
	if err != nil {
		fmt.Println(err)
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	mb.MetricsConfig = &RawMetricsConfig{}
	err = cfgfile.Read(mb.MetricsConfig, "")
	if err != nil {
		fmt.Println(err)
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	logp.Info("Setup base and raw configuration for Modules and Metrics")
	// Apply the base configuration to each module and metric
	for moduleName, module := range helper.Registry {
		// Check if config for module exist
		if _, ok := mb.MbConfig.Metricbeat.Modules[moduleName]; !ok {
			continue;
		}
		module.BaseConfig = mb.MbConfig.Metricbeat.Modules[moduleName]
		module.RawConfig = mb.ModulesConfig.Metricbeat.Modules[moduleName]
		module.Enabled = true

		for metricName, metric := range module.Metrics {

			if _, ok := mb.MbConfig.Metricbeat.Modules[moduleName].Metrics[metricName]; !ok {
				continue;
			}
			metric.BaseConfig = mb.MbConfig.Metricbeat.Modules[moduleName].Metrics[metricName]
			metric.RawConfig = mb.MetricsConfig.Metricbeat.Modules[moduleName].Metrics[metricName]
			metric.Enabled = true
		}
	}

	return nil
}

func (mb *Metricbeat) Setup(b *beat.Beat) error {
	mb.done = make(chan struct{})
	return nil
}

func (mb *Metricbeat) Run(b *beat.Beat) error {
	var err error

	ListAll()

	helper.StartModules(b)

	for {

		select {
		case <-mb.done:
			return nil
		}
	}

	return err
}

func (mb *Metricbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (mb *Metricbeat) Stop() {
	close(mb.done)
}
