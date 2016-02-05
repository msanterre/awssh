package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdConnect = &Command{
	Usage:     "connect",
	Short:     "Connect to an instance",
	Run:       runConnect,
	Shortname: "c",
}

func (machine *Machine) sshString() string {
	return machine.User + "@" + machine.Address
}

func runConnect(cmd *Command, args []string) {
	createStorageIfNotExists()

	if len(args) == 1 {
		file := machineFile(args[0])
		machine := getMachine(file)
		connStr := machine.sshString()

		if machine.User == "" || machine.Address == "" {
			fmt.Println("Could not find an instance for:", args[0])
		} else {
			fmt.Println("Connecting to:", connStr)

			cmd := exec.Command("ssh", connStr)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Run()
		}
	} else {
		fmt.Println("Usage: awssh connect [instance]")
	}
}
