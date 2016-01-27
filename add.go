package main

var cmdAdd = &Command{
	Usage: "add",
	Short: "add a new machine",
	Run:   runAdd,
}

func runAdd(cmd *Command, args []string) {
	createStorageIfNotExists()

}
