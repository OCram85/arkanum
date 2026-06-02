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
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
