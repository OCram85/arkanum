package session

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitea.ocram85.com/arkanum/arkanum/internal/logger"
	"gitea.ocram85.com/arkanum/arkanum/internal/shell"
)


func sessionFile() string {
	return filepath.Join(shell.HomeDir(), "arkanum-session")
}

// Save appends the given items to the session file.
func Save(_ context.Context, items []string) error {
	if len(items) == 0 {
		return fmt.Errorf("no items provided to save")
	}
	f, err := os.OpenFile(sessionFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("session save: %w", err)
	}
	defer f.Close()
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		logger.Info(fmt.Sprintf("Adding '%s' to your session config", item))
		fmt.Fprintln(f, item)
	}
	return nil
}

// Restore reads the session file and calls installFn for each entry.
// installFn is injected from cmd to avoid circular imports.
func Restore(c context.Context, installFn func(context.Context, string) error) error {
	path := sessionFile()
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("there is no arkanum session")
		}
		return fmt.Errorf("session restore: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		item := strings.TrimSpace(scanner.Text())
		if item == "" {
			continue
		}
		logger.Info(fmt.Sprintf("Restoring '%s'...", item))
		if err := installFn(c, item); err != nil {
			logger.Error(fmt.Sprintf("failed to restore '%s': %v", item, err))
		}
	}
	logger.Info("Arkanum session restore completed. 🏁")
	return scanner.Err()
}

// Show prints the current session file contents.
func Show(_ context.Context) error {
	path := sessionFile()
	logger.Info(fmt.Sprintf("Trying to read your session config '%s':", path))
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("there is no session defined")
		}
		return fmt.Errorf("session show: %w", err)
	}
	fmt.Print(string(data))
	return nil
}

// Reset deletes the session file.
func Reset(_ context.Context) error {
	path := sessionFile()
	logger.Info(fmt.Sprintf("Deleting session file '%s'...", path))
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("session reset: %w", err)
	}
	return nil
}
