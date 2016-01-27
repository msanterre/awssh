package main

import (
  os
)

var cmdRemove = &Command{
	Usage: "remove",
	Short: "remove a machine",
	Run:   runRemove,
}

func machineFile(name string) bool {
	machinesDir := os.ExpandEnv(StorageDest)
  filePath := path.Join(
}

func runRemove(cmd *Command, args []string) {
	createStorageIfNotExists()

}
