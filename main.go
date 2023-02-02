/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"
	"strings"

	"github.com/aak1247/pathfit/cmd"
)

func getEnvs() {
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			println(string(parts[0]), string(parts[1]))
		}
	}
}

func main() {
	// getEnvs()
	cmd.Execute()
}
