package main

import (
	"fmt"
	"os"

	"github.com/joeymiller/trooper/console"
	"github.com/joeymiller/trooper/server"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("trooper", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"console": func() (cli.Command, error) {
			return &console.ConsoleCommands{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorBlue,
					ErrorColor:  cli.UiColorRed,
				},
			}, nil
		},
		"server": func() (cli.Command, error) {
			return &server.ServerCommands{
				Ui: &cli.ColoredUi{
					Ui:          ui,
					OutputColor: cli.UiColorBlue,
					ErrorColor:  cli.UiColorRed,
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
