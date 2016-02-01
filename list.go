package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var cmdList = &Command{
	Usage: "list",
	Short: "List the saves instances",
	Run:   runList,
}

func getFiles() []os.FileInfo {
	machinesDir := os.ExpandEnv(StorageDest)
	fileInfos, _ := ioutil.ReadDir(machinesDir)

	return fileInfos
}

func writeHeaders() {
	fmt.Printf("%-20s %-10s %s\n", " Name", " User", " Address")
	fmt.Printf("%-20s %-10s %s\n", "-----------------", "------", "---------")
}

func writeMachine(machine *Machine) {
	fmt.Printf("%-20s %-10s %s\n", machine.Name, machine.User, machine.Address)
}

func runList(cmd *Command, args []string) {
	createStorageIfNotExists()
	fileInfos := getFiles()

  if len(fileInfos) > 0 {
    for _, fileInfo := range fileInfos {
      machinesDir := os.ExpandEnv(StorageDest)
      filePath := path.Join(machinesDir, fileInfo.Name())
      machine := getMachine(filePath)
      writeMachine(machine)
    }
  } else {
    fmt.Println("You don't have any instances yet!")
  }
}
