package jobs

import (
	"context"
	"fmt"

	"github.com/adorigi/localstack-deployment/pkg/terraform"
	"github.com/adorigi/workerpool"
	"github.com/google/uuid"
)

type TerraformJob struct {
	// resourceType       string
	terraformObject    terraform.TerrformObject
	tfvarsJsonFilePath string
	backendConfig      []string
	// workerPool         *workerpool.WorkerPool
	// AzureADCredentials azure.AzureADCredentials
	workerpool.TaskProperties
}

func NewTerraformJob(
	// resourceType string,
	workingDirectory string,
	execPath string,
	tfvarsJsonFilePath string,
	backendConfig []string,
	// workerPool *workerpool.WorkerPool,
	// azureCredentials azure.AzureADCredentials,
) *TerraformJob {
	return &TerraformJob{
		// resourceType: resourceType,
		TaskProperties: workerpool.TaskProperties{
			ID:          uuid.New(),
			Description: "Provisioning",
		},
		terraformObject:    terraform.NewTerraformObject(workingDirectory, execPath),
		tfvarsJsonFilePath: tfvarsJsonFilePath,
		backendConfig:      backendConfig,
		// workerPool:         workerPool,
		// AzureADCredentials: azureCredentials,
	}
}

func (tj *TerraformJob) Properties() workerpool.TaskProperties {
	return tj.TaskProperties
}

func (tj *TerraformJob) Run(ctx context.Context) error {

	planFileName := fmt.Sprintf("%s.tfplan", tj.TaskProperties.ID.String())

	err := tj.terraformObject.Initialize()
	if err != nil {
		return fmt.Errorf("initialize failed: %s %w", tj.TaskProperties.ID.String(), err)
	}
	err = tj.terraformObject.Init(ctx, tj.backendConfig)
	if err != nil {
		return fmt.Errorf("init failed: %w", err)
	}

	err = tj.terraformObject.Plan(ctx, tj.tfvarsJsonFilePath, planFileName)
	if err != nil {
		return fmt.Errorf("plan failed: %w", err)
	}

	err = tj.terraformObject.Apply(ctx, planFileName)
	if err != nil {
		return fmt.Errorf("apply failed: %w", err)
	}

	return nil

}
