package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Config arkanum itself",
				EnableShellCompletion: true,
				Commands: []*cli.Command{
					{
						Name:   "disable-motd",
						Usage:  "Disables hint in new bash terminal",
						Action: disableMotd,
					},
					{
						Name:   "install-extensions",
						Usage:  "Installs predefined recommended  extensions",
						Action: installDefaultExtensions,
					},
					{
						Name:   "reset-codesettings",
						Usage:  "Sets VS Code user setting with basic (Fira Code)",
						Action: setCodeSettings,
					},
				},
			},
			{
				Name: "git",
				Aliases: []string{"g"},
				Usage: "The git command is a wrapper for git helpers",
				Commands: []*cli.Command{
					{
						Name: "setup",
						Usage: "Takes user name and email address to setup the git client",
						Action: SetGitConfig,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
