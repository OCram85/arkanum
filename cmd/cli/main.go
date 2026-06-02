package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		EnableShellCompletion: true,
		Name:                  "arkanum",
		Version:               "2.0.0",
		Usage:                 "✨🌌☄️💥 Install optional tools for developing in a code-server container environment",
		Action: func(ctx context.Context, c *cli.Command) error {
			return cli.ShowRootCommandHelp(c)
		},
		ExitErrHandler: func(ctx context.Context, c *cli.Command, err error) {
			ExitOnErr(err)
		},
		Commands: []*cli.Command{
			ConfigCmd(),
			GitCmd(),
			InstallCmd(),
			SessionCmd(),
		},
		/*CustomRootCommandHelpTemplate: `NAME:
		   {{.FullName}} - {{.Usage}}

		USAGE:
		   {{.FullName}} [command] [subcommand] [arguments]

		VERSION:
		   {{.Version}}

		COMMANDS:
		   {{range .Commands}}{{.Name}}{{"\t"}}{{.Usage}}
		   {{end}}
		EXAMPLES:
		   arkanum git setup "my-name" "my-email"
		   arkanum install golang
		   arkanum install golang 1.22.0
		   arkanum config disable-motd
		   arkanum session save lazygit powershell gitea

		Run '{{.FullName}} [command] --help' to see subcommands.
		`,*/
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
