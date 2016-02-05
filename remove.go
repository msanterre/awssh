package main

import (
	"fmt"
	"os"
)

var cmdRemove = &Command{
	Usage:     "remove",
	Short:     "Remove an instance",
	Run:       runRemove,
	Shortname: "r",
}

func runRemove(cmd *Command, args []string) {
	createStorageIfNotExists()

	for _, machine := range args {
		machinePath := machineFile(machine)
		err := os.Remove(machinePath)

		if err == nil {
			fmt.Println("Removed instance:", machine)
		} else {
			fmt.Println("No instance found for:", machine)
		}
	}

}
