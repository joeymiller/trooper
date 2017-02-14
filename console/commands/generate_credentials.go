package commands

import (
	"bytes"
	"flag"
	"text/template"

	"github.com/joeymiller/creddie/credentials"
	"github.com/mitchellh/cli"
)

const consoleTemplate = `
Access Key ID: {{.AccessKeyId}}
Secret Access Key: {{.SecretAccessKey}}
Session Token: {{.SessionToken}}
Expires: {{.Expiration}}
`

type GenerateCredentialsCommand struct {
	Ui   cli.Ui
	Role string
}

func (g *GenerateCredentialsCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("console", flag.ContinueOnError)
	cmdFlags.StringVar(&g.Role, "role", "playground", "The role assumed by the credentials.")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	resp, err := requestcredentials.Generate(g.Role)
	if err != nil {
		g.Ui.Error("Access Denied!")
		return 1
	}

	temp := template.Must(template.New("ConsoleOutput").Parse(consoleTemplate))
	buf := &bytes.Buffer{}
	if err := temp.Execute(buf, resp.Credentials); err != nil {
		g.Ui.Error("Error! Please Try Again.")
		return 1
	}

	g.Ui.Output(buf.String())
	return 0
}

func (g *GenerateCredentialsCommand) Help() string {
	return "Returns a set of temporary security credentials (consisting of an access key ID, a secret access key, and a security token) that you can use to access AWS resources"
}

func (g *GenerateCredentialsCommand) Synopsis() string {
	return "Returns a set of temporary security credentials"
}
