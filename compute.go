package main

import (
	"github.com/codegangsta/cli"
)

// getComputeCmds return commands for compute section
func getComputeCmds() (computeCmds []cli.Command) {

	// Ip commands
	computeCmds = []cli.Command{
		// getProperties
		{
			Name:        "listFlavors",
			Usage:       "Get compute instance flavors",
			Description: "ra compute listFlavors REGION",
			Action: func(c *cli.Context) {
				dieOk()
			},
		},
	}
	return

}
