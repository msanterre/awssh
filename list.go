package main

import (
	"fmt"
	"io/ioutil"
	"os"
  "path"
  "encoding/json"
)

var cmdList = &Command{
	Usage: "list",
	Short: "list the saves machines",
	Run:   runList,
}

func runList() {
	createStorageIfNotExists()
	machinesDir := os.ExpandEnv(StorageDest)
	fileInfos, _ := ioutil.ReadDir(machinesDir)

  fmt.Printf("%-20s %-10s %s\n", " Name", " User", " Address")
  fmt.Printf("%-20s %-10s %s\n", "-----------------", "------", "---------")
	for _, fileInfo := range fileInfos {
    filePath := path.Join(machinesDir, fileInfo.Name())
    fileContent, _ := ioutil.ReadFile(filePath)
    machine := &Machine{}
    json.Unmarshal(fileContent, &machine)
    fmt.Printf("%-20s %-10s %s\n", machine.Name, machine.User, machine.Address)
	}
}
