package gitcfg

import (
	"context"
	"fmt"

	"gitea.ocram85.com/arkanum/arkanum/internal/logger"
	"gitea.ocram85.com/arkanum/arkanum/internal/shell"
)


// Setup configures the global git user.name and user.email.
func Setup(c context.Context, username, email string) error {
	if username == "" {
		return fmt.Errorf("invalid or empty username given")
	}
	if email == "" {
		return fmt.Errorf("invalid or empty email given")
	}
	logger.Info("Setting global git config...")
	if err := shell.ExecCommand(c, "git", "config", "--global", "user.name", username); err != nil {
		return fmt.Errorf("git config user.name: %w", err)
	}
	if err := shell.ExecCommand(c, "git", "config", "--global", "user.email", email); err != nil {
		return fmt.Errorf("git config user.email: %w", err)
	}
	logger.Info("Returning global config:")
	return shell.ExecCommand(c, "git", "config", "--list", "--global")
}
