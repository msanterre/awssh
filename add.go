package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cmdAdd = &Command{
	Usage: "add",
	Short: "Add a new instance",
	Run:   runAdd,
}

func runAdd(cmd *Command, args []string) {
	createStorageIfNotExists()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adding a new instance:\n")

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Address: ")
	address, _ := reader.ReadString('\n')
	fmt.Print("User: ")
	user, _ := reader.ReadString('\n')

	machine := &Machine{
		Address: strings.TrimSpace(address),
		Name:    strings.TrimSpace(name),
		User:    strings.TrimSpace(user),
	}

	machine.Save()
}
