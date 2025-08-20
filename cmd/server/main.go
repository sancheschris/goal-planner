package main

import (
	"fmt"

	"github.com/sancheschris/goal-planner/configs"
)

func main() {
	// starting point
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(configs)
}

