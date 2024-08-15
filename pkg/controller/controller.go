package controller

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	"github.com/adorigi/localstack-deployment/pkg/config"
	"github.com/adorigi/localstack-deployment/pkg/jobs"
	"github.com/adorigi/workerpool"
)

func Controller() error {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		return fmt.Errorf("cannot find 'terraform' path")
	}

	config, err := config.ParseConfiguration("/Users/adnangulegulzar/GITHUB/adorigi/localstack-deployment/pkg/config/testconfig.json")
	if err != nil {
		return err
	}

	workerPool := workerpool.NewWorkerPool(config.ConcurrentWorkers)
	workerPool.Start(context.Background())

	// for _, parameter := range parameters {

	workingDirectory := path.Join(config.TemplatesDirectory, config.Resource)
	terraformJob := jobs.NewTerraformJob(
		// parameter.ResourceType,
		workingDirectory,
		execPath,
		config.TfvarsJsonFilePath,
		config.BackendConfig,
		// parameter.Vars,
		// workerPool,
		// parameter.Credentials,
	)

	workerPool.AddTask(terraformJob)

	// }

	workerPool.Wait()

	return nil

}
