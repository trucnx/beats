package cfgfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

// Command line flags
var configfile *string
var testConfig *bool

func init() {
	// The default config cannot include the beat name as it is not initialised when this
	// function is called, but see ChangeDefaultCfgfileFlag
	configfile = flag.String("c", "/etc/beat/beat.yml", "Configuration file")
	testConfig = flag.Bool("configtest", false, "Test configuration and exit.")
}

// ChangeDefaultCfgfileFlag replaces the value and default value for the `-c` flag so that
// it reflects the beat name.
func ChangeDefaultCfgfileFlag(beatName string) error {
	cliflag := flag.Lookup("c")
	if cliflag == nil {
		return fmt.Errorf("Flag -c not found")
	}

	if runtime.GOOS == "windows" {
		cliflag.DefValue = fmt.Sprintf(`C:\Program Files\%s\%s.yml`,
			strings.Title(beatName), beatName)
	} else {
		cliflag.DefValue = fmt.Sprintf("/etc/%s/%s.yml", beatName, beatName)
	}
	return cliflag.Value.Set(cliflag.DefValue)
}

// Read reads the configuration from a yaml file into the given interface structure.
// In case path is not set this method reads from the default configuration file for the beat.
func ReadConfigFile(out interface{}, path string) error {

	filecontent, err := ioutil.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Failed to read %s: %v. Exiting.", path, err)
	}

	err = loadConfig(out, filecontent)
	if (err != nil) {
		return fmt.Errorf("Configuration error: %v. Exiting.", err)
	}

	return nil
}

func IsTestConfig() bool {
	return *testConfig
}

// Loads the given YAML config string into the out interface
func loadConfig(config interface{}, yamlString []byte) error {
	if err := yaml.Unmarshal(yamlString, config); err != nil {
		return fmt.Errorf("YAML config parsing failed: %v. Exiting.", err)
	}

	return nil
}

//func convertToJson(yamlConfig []byte) []byte {
//
//}
//
//func convertToYaml(jsonConfig []byte) []byte {
//
//}
