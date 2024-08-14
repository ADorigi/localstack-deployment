package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

// basically perform 'terraform init'
func (to *TerrformObject) Init(ctx context.Context, backendConfig []string) error {

	initOpitons := []tfexec.InitOption{
		tfexec.Upgrade(true),
	}

	for _, backendOption := range backendConfig {
		initOpitons = append(initOpitons, tfexec.BackendConfig(backendOption))
	}

	err := to.tf.Init(ctx, initOpitons...)
	if err != nil {
		return err
	}
	return nil

}
