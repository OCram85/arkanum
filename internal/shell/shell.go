package shell

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// ExecCommand runs a command, streaming stdout/stderr to the terminal.
func ExecCommand(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// ExecSudoCommand runs a command via `sudo -E` to preserve env vars (e.g. proxy).
// Set hideStdout to true to suppress output (e.g. for apt-get update noise).
func ExecSudoCommand(ctx context.Context, hideStdout bool, name string, arg ...string) error {
	args := append([]string{"-E", name}, arg...)
	cmd := exec.CommandContext(ctx, "sudo", args...)
	if hideStdout {
		cmd.Stdout = nil
	} else {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// HomeDir returns the current user's home directory.
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not determine home directory: %v\n", err)
		os.Exit(1)
	}
	return home
}
