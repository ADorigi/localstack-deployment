package main

import (
	"github.com/adorigi/localstack-deployment/pkg/controller"
)

func main() {

	err := controller.Controller()
	if err != nil {
		panic(err)
	}
}
