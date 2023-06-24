package main

import (
	"fmt"

	"github.com/mr-destructive/palm"
)

func main() {

	models, err := palm.ListModels()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(models)
	modelName := models[0].Name
	fmt.Println(modelName)
	model, err := palm.GetModel(modelName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(model)

}
