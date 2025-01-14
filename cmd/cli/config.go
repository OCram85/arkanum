package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func disableMotd(ctx context.Context, cmd *cli.Command) error {
	say("Disabling 'arkanum' motd...")
	homeDir := os.Getenv("HOME")
	motdFile := filepath.Join(homeDir, "enable_motd")
	if _, err := os.Stat(motdFile); err != nil {
		sayWarning("Arkanum Motd already disabled")
		return nil
	}

	err := os.Remove(motdFile)
	if err != nil {
		return err
	}
	return nil
}

func installDefaultExtensions(ctx context.Context, cmd *cli.Command) error {
	say("Installing default extensions...")
	// Gitlens
	if err := installSingleExtension(ctx, "eamodio.gitlens"); err != nil {
		return nil
	}
	if err := installSingleExtension(ctx, "zhuangtongfa.material-theme"); err != nil {
		return nil
	}
	if err := installSingleExtension(ctx, "vscode-icons-team.vscode-icons"); err != nil {
		return nil
	}
	say("done.")
	return nil
}

func getDefaultSettings() string {
	tmp := `{
   "window.menuBarVisibility": "compact",
  "workbench.colorTheme": "One Dark Pro Darker",
  "workbench.iconTheme": "vscode-icons",
  "editor.fontFamily": "'FiraCode', 'FiraCode Nerd Font', 'FiraCode NF', Consolas, 'Courier New', monospace",
  "terminal.integrated.fontFamily": "'FiraCode Mono', 'FiraCode Nerd Font Mono', 'FiraCode NFM', Consolas, monospace",
  "editor.fontLigatures": true,
  "editor.formatOnSave": true,
  "extensions.autoUpdate": false,
  "git.confirmSync": false,
  "telemetry.telemetryLevel": "off"
}`
	return tmp
}

func setCodeSettings(ctx context.Context, cmd *cli.Command) error {
	home := os.Getenv("HOME")
	userSettings := path.Join(home, "data/User/settings.json")
	defaults := getDefaultSettings()
	if _, err := os.Stat(userSettings); err == nil {
		sayWarning("Settings file found. Resetting content...")
		// No need to remove file WriteFile dumps content
		//if err := os.Remove(userSettings); err != nil {
		//	return err
		//}
	}
	say(fmt.Sprintf("Writing Arkanum base settings.(%v)", userSettings))
	if err := os.WriteFile(userSettings, []byte(defaults), 0644); err != nil {
		return err
	}
	say("done.")
	return nil
}
