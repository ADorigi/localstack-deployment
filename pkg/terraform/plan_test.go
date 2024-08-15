package terraform

import (
	"context"
	"os/exec"
	"testing"
)

func TestPlan(t *testing.T) {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Errorf("Cannot find 'terraform' path")
	}

	to := NewTerraformObject("templates/eks", execPath)

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

	err = to.Plan(context.Background(), "/Users/adnangulegulzar/GITHUB/adorigi/localstack-deployment/pkg/terraform/testdata/plan.tfvars.json", "test.tfplan")
	if err != nil {
		t.Log(err)
		t.Errorf("Cannot create terraform plan")
	}

}
