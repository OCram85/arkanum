package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func SetGitConfig(ctx context.Context, cmd *cli.Command) error {
	say("Setting global git config...")
	if cmd.Args().Len() != 2 {
		sayError(fmt.Sprintf("Unknown arguments count (%v)", cmd.Args().Len()))
	}
	//err := execCommand(ctx, "git", "--global", "user.name", cmd.Args().Get(0))
	return nil
}
