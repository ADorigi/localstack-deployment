package config

import (
	"encoding/json"
	"io"
	"os"
)

type ControllerConfiguration struct {
	ConcurrentWorkers int      `json:"concurrent_workers"`
	WorkingDirectory  string   `json:"working_directory"`
	Vars              []string `json:"vars"`
}

// var ConcurrentWorkers = 1
// var WorkingDirectory = "pkg/terraform/templates"

func ParseConfiguration(configJsonFile string) ([]*ControllerConfiguration, error) {

	var params []*ControllerConfiguration

	jsonFile, err := os.Open(configJsonFile)
	if err != nil {
		return nil, err
	}

	byteArray, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(byteArray, &params); err != nil {
		return nil, err
	}

	return params, nil

}
