package main

import (
	"os"
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
