package main

import (
  "os"
  "strings"
  "bufio"
  "fmt"
)

var cmdClear = &Command{
	Usage: "clear",
	Short: "Clear all instances",
	Run:   runClear,
}


func runClear(cmd *Command, args []string) {
	createStorageIfNotExists()

  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Are you sure you want to clear all instances? (y/n):")

  confirm, _ := reader.ReadString('\n')

  if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
    machinesDir := os.ExpandEnv(StorageDest)
    os.RemoveAll(machinesDir)
    fmt.Println("Instances cleared!")
  }

}
