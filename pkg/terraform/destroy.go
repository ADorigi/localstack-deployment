package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func (to *TerrformObject) Destroy(ctx context.Context, tfvarsJsonFile string) error {

	destroyOptions := []tfexec.DestroyOption{
		tfexec.VarFile(tfvarsJsonFile),
	}

	// for _, tfvar := range tfvars {
	// 	destroyOptions = append(destroyOptions, tfexec.Var(tfvar))
	// }

	err := to.tf.Destroy(ctx, destroyOptions...)
	if err != nil {
		return err
	}
	return nil
}
