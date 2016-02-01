package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const (
	DefaultUser = "ubuntu"
	StorageDest = "$HOME/.awssh/machines"
)

type Machine struct {
	Name    string `json:"name"`
	User    string `json:"user"`
	Address string `json:"host"`
}

func getMachine(filePath string) *Machine {
	fileContent, _ := ioutil.ReadFile(filePath)
	machine := &Machine{}
	json.Unmarshal(fileContent, &machine)

	return machine
}

func machineFile(name string) string {
	machinesDir := os.ExpandEnv(StorageDest)
	return path.Join(machinesDir, name)
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createStorageIfNotExists() {
	expandedDest := os.ExpandEnv(StorageDest)
	storageExists, err := fileExists(expandedDest)
	fail(err)

	if !storageExists {
		err = os.MkdirAll(expandedDest, 0777)
		fail(err)
	}
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
