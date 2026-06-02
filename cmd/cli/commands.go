package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"gitea.ocram85.com/arkanum/arkanum/internal/config"
	"gitea.ocram85.com/arkanum/arkanum/internal/gitcfg"
	"gitea.ocram85.com/arkanum/arkanum/internal/installer"
	"gitea.ocram85.com/arkanum/arkanum/internal/logger"
	"gitea.ocram85.com/arkanum/arkanum/internal/session"
)

// installByName dispatches a tool name to its installer function.
// Used by both the install command and session restore.
func installByName(ctx context.Context, tool string) error {
	switch tool {
	case "docker-cli":
		return installer.DockerCLI(ctx)
	case "dotnet":
		return installer.DotNet(ctx)
	case "golang":
		return installer.GoLang(ctx, "")
	case "bun":
		return installer.Bun(ctx)
	case "nodejs":
		return installer.NodeJS(ctx)
	case "volta":
		return installer.Volta(ctx)
	case "powershell":
		return installer.PowerShell(ctx)
	case "gitea":
		return installer.GiteaTools(ctx)
	case "lazygit":
		return installer.LazyGit(ctx)
	default:
		return fmt.Errorf("unknown tool: %q", tool)
	}
}

// ConfigCmd returns the `config` command with all its sub-commands.
func ConfigCmd() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Modify arkanum itself",
		Commands: []*cli.Command{
			{
				Name:  "disable-motd",
				Usage: "Disable the hint shown in new bash terminals",
				Action: func(ctx context.Context, c *cli.Command) error {
					return config.DisableMotd(ctx)
				},
			},
			{
				Name:  "install-extensions",
				Usage: "Install predefined recommended VS Code extensions",
				Action: func(ctx context.Context, c *cli.Command) error {
					return config.InstallExtensions(ctx)
				},
			},
			{
				Name:  "reset-codesettings",
				Usage: "Set VS Code user settings with basic Fira Code config",
				Action: func(ctx context.Context, c *cli.Command) error {
					return config.ResetCodeSettings(ctx)
				},
			},
		},
	}
}

// GitCmd returns the `git` command with all its sub-commands.
func GitCmd() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "Git helpers",
		Commands: []*cli.Command{
			{
				Name:      "setup",
				Usage:     "Configure global git user.name and user.email",
				ArgsUsage: "<username> <email>",
				Action: func(ctx context.Context, c *cli.Command) error {
					if c.NArg() < 2 {
						return fmt.Errorf("please provide <username> and <email>")
					}
					return gitcfg.Setup(ctx, c.Args().Get(0), c.Args().Get(1))
				},
			},
		},
	}
}

// InstallCmd returns the `install` command with all its sub-commands.
func InstallCmd() *cli.Command {
	return &cli.Command{
		Name:  "install",
		Usage: "Install optional tools",
		Commands: []*cli.Command{
			{
				Name:   "docker-cli",
				Usage:  "Install the latest docker-cli",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.DockerCLI(ctx) },
			},
			{
				Name:   "dotnet",
				Usage:  "Install latest LTS .NET Core SDK + runtime",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.DotNet(ctx) },
			},
			{
				Name:      "golang",
				Usage:     "Install latest (or specified) Go version",
				ArgsUsage: "[version]",
				Action: func(ctx context.Context, c *cli.Command) error {
					return installer.GoLang(ctx, c.Args().First())
				},
			},
			{
				Name:   "bun",
				Usage:  "Install latest Bun version",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.Bun(ctx) },
			},
			{
				Name:   "nodejs",
				Usage:  "Install latest NodeJS LTS via Volta",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.NodeJS(ctx) },
			},
			{
				Name:   "volta",
				Usage:  "Install Volta as NodeJS version manager",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.Volta(ctx) },
			},
			{
				Name:   "powershell",
				Usage:  "Install latest PowerShell LTS",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.PowerShell(ctx) },
			},
			{
				Name:   "gitea",
				Usage:  "Install Gitea tools (tea + changelog)",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.GiteaTools(ctx) },
			},
			{
				Name:   "lazygit",
				Usage:  "Install latest Lazygit binary",
				Action: func(ctx context.Context, c *cli.Command) error { return installer.LazyGit(ctx) },
			},
			{
				Name: "just",
				Usage: "Install just. a modern makefile alternative",
				Action: func(ctx context.Context, c *cli.Command) error {return installer.Just(ctx)},
			},
		},
	}
}

// SessionCmd returns the `session` command with all its sub-commands.
func SessionCmd() *cli.Command {
	return &cli.Command{
		Name:  "session",
		Usage: "Manage install sessions",
		Commands: []*cli.Command{
			{
				Name:      "save",
				Usage:     "Add install items to the session config",
				ArgsUsage: "<tool> [tool...]",
				Action: func(ctx context.Context, c *cli.Command) error {
					if c.NArg() == 0 {
						return fmt.Errorf("please provide at least one tool name")
					}
					return session.Save(ctx, c.Args().Slice())
				},
			},
			{
				Name:  "restore",
				Usage: "Restore and install all items from the saved session",
				Action: func(ctx context.Context, c *cli.Command) error {
					return session.Restore(ctx, installByName)
				},
			},
			{
				Name:  "show",
				Usage: "Show the currently defined session content",
				Action: func(ctx context.Context, c *cli.Command) error {
					return session.Show(ctx)
				},
			},
			{
				Name:  "reset",
				Usage: "Delete the session config",
				Action: func(ctx context.Context, c *cli.Command) error {
					return session.Reset(ctx)
				},
			},
		},
	}
}

// ExitOnErr logs an error and exits with code 1.
func ExitOnErr(err error) {
	if err != nil {
		logger.Error(err.Error())
		cli.OsExiter(1)
	}
}
