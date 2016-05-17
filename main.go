package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// Usage is the one-line usage message.
	// The first word in the line is taken to be the command name.
	Usage string

	// Short name is a shorthand for the long name
	Shortname string

	// Short is the short description shown in the 'awssh help' output.
	Short string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

var commands = []*Command{
	cmdSync,
	cmdList,
	cmdAdd,
	cmdRemove,
	cmdClear,
	cmdConnect,
}

func main() {
	flag.Usage = UsageExit
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("awssh: ")
	args := flag.Args()
	if len(args) < 1 {
		UsageExit()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] || cmd.Shortname == args[0] {
			cmd.Flag.Usage = func() { cmd.UsageExit() }
			cmd.Flag.Parse(args[1:])
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}

	fmt.Fprintf(os.Stderr, "awssh: unknown command %q\n", args[0])
	fmt.Fprintf(os.Stderr, "Run 'awssh help' for usage.\n")
	os.Exit(2)
}

func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: awssh help command\n\n")
		fmt.Fprintf(os.Stderr, "Too many arguments given.\n")
		os.Exit(2)
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			tmpl(os.Stdout, helpTemplate, cmd)
			return
		}
	}
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	})
	template.Must(t.Parse(strings.TrimSpace(text) + "\n\n"))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

var usageTemplate = `
Awssh is a small tool for easily managing and using your ec2 host connections

Usage:

    awssh command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-8s"}} {{.Short}}{{end}}

`

var helpTemplate = `
Usage: awssh {{.Usage}}

{{.Short | trim}}
`

func UsageExit() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func (c *Command) UsageExit() {
	fmt.Fprintf(os.Stderr, "Usage: awssh %s\n\n", c.Usage)
	fmt.Fprintf(os.Stderr, "Run 'awssh help %s' for help.\n", c.Name())
	os.Exit(2)
}
