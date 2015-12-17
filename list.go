package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var cmdList = &Command{
	Usage: "list",
	Short: "list the saves machines",
	Run:   runList,
}

func getFiles() []os.FileInfo {
	machinesDir := os.ExpandEnv(StorageDest)
	fileInfos, _ := ioutil.ReadDir(machinesDir)

	return fileInfos
}

func getMachine(fileInfo os.FileInfo) *Machine {
  machinesDir := os.ExpandEnv(StorageDest)
	filePath := path.Join(machinesDir, fileInfo.Name())
	fileContent, _ := ioutil.ReadFile(filePath)
	machine := &Machine{}
	json.Unmarshal(fileContent, &machine)

  return machine
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

	for _, fileInfo := range fileInfos {
		machine := getMachine(fileInfo)
		writeMachine(machine)
	}
}
