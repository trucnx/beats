package main

import (
beat "github.com/elastic/beats/libbeat/beat"
	filebeat "github.com/elastic/beats/filebeat/beat"
	topbeat "github.com/elastic/beats/topbeat/beat"
	packetbeat "github.com/elastic/beats/packetbeat/beat"
	"io/ioutil"
	"time"
	"fmt"
	"github.com/elastic/beats/libbeat/logp"
	"os"
	"flag"
)

var Version = "1.0.0"
var Name = "managerbeat"

func main() {



	mb := &Managerbeat{}
	mb.Run()
}

type Managerbeat struct {
	done chan struct{}
}


func (t *Managerbeat) Run() error {

	fmt.Print("Managebeat started")
	StartBeats()
	d, _ := time.ParseDuration("1s")
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-t.done:
			return nil
		case <-ticker.C:
		}
	}

	return nil
}

func StartBeats() {



	func go() {
		filebeatconfig, err := ioutil.ReadFile("../filebeat/etc/filebeat.yml")
		if err != nil {
			fmt.Println(err)
		}
		RunBeat("filebeat", "1.0.0", filebeat.New(), filebeatconfig)
	}()



	func go() {

		packetbeatconfig, err := ioutil.ReadFile("../packetbeat/etc/packetbeat.yml")
		if err != nil {
			fmt.Println(err)
		}

		RunBeat("packetbeat", "1.0.0", packetbeat.New(), packetbeatconfig)
	}()

	func go() {

		topbeatconfig, err := ioutil.ReadFile("../topbeat/etc/topbeat.yml")
		if err != nil {
			fmt.Println(err)
		}

		RunBeat("topbeat", "1.0.0", topbeat.New(), topbeatconfig)
	}()

}


// Initiates and runs a new beat object
func RunBeat(name string, version string, bt beat.Beater, config string) *beat.Beat {
	b := beat.NewBeat(name, version, bt)

	// Additional command line args are used to overwrite config options
	b.CommandLineSetup()

	flag.Set("config-string", config)

	// Loads base config
	b.LoadConfig()

	// Configures beat
	err = bt.Config(b)
	if err != nil {
		logp.Critical("Config error: %v", err)
		os.Exit(1)
	}

	// Run beat. This calls first beater.Setup,
	// then beater.Run and beater.Cleanup in the end
	b.Run()

	return b
}
