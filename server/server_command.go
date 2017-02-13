package server

import (
	"github.com/joeymiller/trooper/server/commands"
	"github.com/mitchellh/cli"
)

type ServerCommands struct {
	Ui cli.Ui
}

func (s *ServerCommands) Run(args []string) int {
	servercmd := cli.NewCLI("trooper server", "")
	servercmd.Args = args

	servercmd.Commands = map[string]cli.CommandFactory{
		"run": func() (cli.Command, error) {
			return &commands.ServerRunCommand{Ui: s.Ui}, nil
		},
	}

	if exitStatus, err := servercmd.Run(); err != nil {
		s.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}
}

func (s *ServerCommands) Help() string {
	return "trooper Server Commands."
}

func (s *ServerCommands) Synopsis() string {
	return "Run trooper as Server."
}
