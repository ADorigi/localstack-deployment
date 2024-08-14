package terraform

import (
	"context"
	"os/exec"
	"testing"
)

func TestInit(t *testing.T) {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Errorf("Cannot find 'terraform' path")
	}

	to := NewTerraformObject("templates/vpc", execPath)

	err = to.Initialize()
	if err != nil {
		t.Log(err)
		t.Errorf("Initilize failed")
	}

	backendConfig := []string{
		"path=./state",
	}

	err = to.Init(context.Background(), backendConfig)
	if err != nil {
		t.Log(err)
		t.Error("Init failed")
	}

}
