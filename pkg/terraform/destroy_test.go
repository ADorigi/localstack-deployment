package terraform

import (
	"context"
	"os/exec"
	"testing"
)

func TestDestroy(t *testing.T) {

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

	// tfvars := []string{
	// 	"resource_group_name=test-resource-group",
	// 	"location=Canada Central",
	// 	"resourceCount=2",
	// }

	err = to.Destroy(context.Background(), "/Users/adnangulegulzar/GITHUB/adorigi/localstack-deployment/pkg/testdata/two_node_group.tfvars.json")
	if err != nil {
		t.Log(err)
	}

}
