package config

import (
	"slices"
	"testing"
)

var testConfigFilePath = "./testconfig.json"

func TestParseConfiguration(t *testing.T) {

	testdata := struct {
		concurrentWorkers  int
		templatesDirectory string
		resource           string
		tfvarsJsonFilePath string
		backend_config     []string
	}{
		concurrentWorkers:  1,
		templatesDirectory: "/Users/adnangulegulzar/GITHUB/adorigi/localstack-deployment/pkg/terraform/templates",
		resource:           "vpc",
		tfvarsJsonFilePath: "/Users/adnangulegulzar/GITHUB/adorigi/localstack-deployment/pkg/terraform/testdata/plan.tfvars.json",
		backend_config: []string{
			"path=./state",
		},
	}

	config, err := ParseConfiguration(testConfigFilePath)
	if err != nil {
		t.Log(err)
	}

	if config.ConcurrentWorkers != testdata.concurrentWorkers {
		t.Errorf("concurrent workers incorrect")
	}

	if config.TemplatesDirectory != testdata.templatesDirectory {
		t.Errorf("templates directory incorrect")
	}

	if config.Resource != testdata.resource {
		t.Errorf("resource incorrect")
	}

	if config.TfvarsJsonFilePath != testdata.tfvarsJsonFilePath {
		t.Errorf("tfvars file path incorrect")
	}

	for _, bConfig := range config.BackendConfig {
		if slices.Contains(testdata.backend_config, bConfig) != true {
			t.Errorf("backend config incorrect")
		}
	}
}
