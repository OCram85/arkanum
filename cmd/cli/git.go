package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func SetGitConfig(ctx context.Context, cmd *cli.Command) error {
	say("Setting global git config...")
	if cmd.Args().Len() != 2 {
		err := sayError(fmt.Sprintf("Unknown arguments count (%v)", cmd.Args().Len()))
		return err
	}
	var err error
	gitUser := cmd.Args().Get(0)
	say(fmt.Sprintf("Setting global git user.name to '%v'...", gitUser))
	err = execCommand(ctx, "git", "config", "--global", "user.name", gitUser)
	if err != nil {
		return err
	}

	gitEmail := cmd.Args().Get(1)
	say(fmt.Sprintf("Setting global git user.email to '%v'...", gitEmail))
	err = execCommand(ctx, "git", "config", "--global", "user.email", cmd.Args().Get(1))
	if err != nil {
		return err
	}
	return nil
}
