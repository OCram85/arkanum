package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gitea.ocram85.com/arkanum/arkanum/internal/logger"
	"gitea.ocram85.com/arkanum/arkanum/internal/shell"
)


// DisableMotd removes the motd enable-flag from the user's home directory.
func DisableMotd(_ context.Context) error {
	flag := filepath.Join(shell.HomeDir(), "enable_motd")
	if _, err := os.Stat(flag); os.IsNotExist(err) {
		logger.Warn("Arkanum Motd already disabled")
		return nil
	}
	logger.Info("Disabling 'arkanum' motd...")
	if err := os.Remove(flag); err != nil {
		return fmt.Errorf("disable-motd: %w", err)
	}
	return nil
}

// InstallExtensions installs the predefined recommended VS Code extensions.
func InstallExtensions(c context.Context) error {
	extensions := []string{
		"eamodio.gitlens",
		"zhuangtongfa.material-theme",
		"vscode-icons-team.vscode-icons",
	}
	names := []string{"gitlens", "One Dark Pro", "vscode-icons"}

	logger.Info("Installing default extensions...")
	for i, ext := range extensions {
		logger.Info(fmt.Sprintf("Installing '%s'...", names[i]))
		if err := shell.ExecCommand(c, "install-extension", ext, "--force"); err != nil {
			return fmt.Errorf("extension %s: %w", ext, err)
		}
	}
	logger.Info("done.")
	return nil
}

// ResetCodeSettings writes the base VS Code user settings file.
func ResetCodeSettings(_ context.Context) error {
	home := shell.HomeDir()
	settingsFile := filepath.Join(home, "data", "User", "settings.json")
	logger.Info(fmt.Sprintf("Setting VSCode base settings (%s)", settingsFile))

	if err := os.MkdirAll(filepath.Dir(settingsFile), 0o755); err != nil {
		return fmt.Errorf("reset-codesettings mkdir: %w", err)
	}

	const settings = `{
  "window.menuBarVisibility": "compact",
  "workbench.colorTheme": "One Dark Pro Darker",
  "workbench.iconTheme": "vscode-icons",
  "editor.fontFamily": "'FiraCode', 'FiraCode Nerd Font', 'FiraCode NF', Consolas, 'Courier New', monospace",
  "terminal.integrated.fontFamily": "'FiraCode Mono', 'FiraCode Nerd Font Mono', 'FiraCode NFM', Consolas, monospace",
  "editor.fontLigatures": true,
  "editor.formatOnSave": true,
  "extensions.autoUpdate": false,
  "git.confirmSync": false,
  "telemetry.telemetryLevel": "off",
  "chat.agent.enabled": false,
  "chat.disableAIFeatures": true,
  "terminal.integrated.initialHint": false
}
`
	if err := os.WriteFile(settingsFile, []byte(settings), 0o644); err != nil {
		return fmt.Errorf("reset-codesettings write: %w", err)
	}
	logger.Info("done.")
	return nil
}
