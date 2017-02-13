package console

import (
	"github.com/joeymiller/trooper/console/commands"
	"github.com/mitchellh/cli"
)

type ConsoleCommands struct {
	Ui cli.Ui
}

func (c *ConsoleCommands) Run(args []string) int {
	console := cli.NewCLI("trooper console", "")
	console.Args = args

	console.Commands = map[string]cli.CommandFactory{
		"generate-credentials": func() (cli.Command, error) {
			return &commands.GenerateCredentialsCommand{Ui: c.Ui}, nil
		},
	}

	if exitStatus, err := console.Run(); err != nil {
		c.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}
}

func (c *ConsoleCommands) Help() string {
	return "trooper Console Commands."
}

func (c *ConsoleCommands) Synopsis() string {
	return "Commands related to trooper Console."
}
