package main

import (
	"fmt"
	"os"
	"slices"
	"time"
	"tree-viz/cli"
	"tree-viz/dir"
)

type CLICommand struct {
	command     string
	description string
	flags       []string
}

var possibleCommands = []CLICommand{
	{
		command:     "dir",
		description: "can be used for tree generation of given path. Example: tree-viz . ./output/",
		flags:       []string{},
	},
}

func getInput() {
	args := os.Args
	viewType := args[1]

	availableInputs := []string{"dir"}

	if !slices.Contains(availableInputs, viewType) {
		invalidCmd(viewType)
		return
	}

	startTime := time.Now()
	switch viewType {
	case "dir":
		dir.GenerateDirTree()
	}
	fmt.Printf("\nTask executed in %s\n", time.Since(startTime))
}

func invalidCmd(invalidInput string) {
	fmt.Printf(cli.Red+"Invalid command '%s'"+cli.Reset+", try using these commands:\n", invalidInput)

	for _, cmd := range possibleCommands {
		fmt.Printf(cli.Green+"    %s"+cli.Reset+": %s "+cli.Cyan+"(flags: %v)"+cli.Reset+"\n", cmd.command, cmd.description, cmd.flags)
	}
}
