package installer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gitea.ocram85.com/arkanum/arkanum/internal/logger"
	"gitea.ocram85.com/arkanum/arkanum/internal/shell"
)

// DockerCLI installs the latest docker-cli.
func DockerCLI(c context.Context) error {
	logger.Info("Installing docker-cli...")

	logger.Info("Getting required packages...")
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("docker-cli apt-get update: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "-y", "ca-certificates", "curl", "gnupg"); err != nil {
		return fmt.Errorf("docker-cli requirements: %w", err)
	}

	logger.Info("Setting up docker repository...")
	if err := shell.ExecSudoCommand(c, false, "install", "-m", "0755", "-d", "/etc/apt/keyrings"); err != nil {
		return fmt.Errorf("docker-cli keyrings dir: %w", err)
	}
	if err := shell.ExecCommand(c, "bash", "-c",
		`curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo -E gpg --dearmor -o /etc/apt/keyrings/docker.gpg`,
	); err != nil {
		return fmt.Errorf("docker-cli gpg: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "chmod", "a+r", "/etc/apt/keyrings/docker.gpg"); err != nil {
		return fmt.Errorf("docker-cli chmod: %w", err)
	}
	if err := shell.ExecCommand(c, "bash", "-c",
		`echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] `+
			`https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | `+
			`sudo -E tee /etc/apt/sources.list.d/docker.list > /dev/null`,
	); err != nil {
		return fmt.Errorf("docker-cli sources.list: %w", err)
	}

	logger.Info("Installing docker-ce-cli package...")
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("docker-cli apt-get update 2: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "--no-install-recommends", "-y", "docker-ce-cli"); err != nil {
		return fmt.Errorf("docker-cli install: %w", err)
	}

	logger.Info("docker-cli done.")
	return nil
}

// DotNet installs the latest LTS .NET Core SDK + runtime.
func DotNet(c context.Context) error {
	logger.Info("Installing dotnet requirements...")
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("dotnet apt-get update: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "--no-install-recommends", "-y", "libicu74"); err != nil {
		return fmt.Errorf("dotnet libicu74: %w", err)
	}

	logger.Info("Downloading latest install script...")
	if err := shell.ExecCommand(c, "curl", "-#fSL", "https://dot.net/v1/dotnet-install.sh", "-o", "/tmp/dotnet-install.sh"); err != nil {
		return fmt.Errorf("dotnet script download: %w", err)
	}
	if err := shell.ExecCommand(c, "chmod", "+x", "/tmp/dotnet-install.sh"); err != nil {
		return fmt.Errorf("dotnet script chmod: %w", err)
	}

	logger.Info("Installing latest .NET Core LTS release...")
	if err := shell.ExecCommand(c, "/tmp/dotnet-install.sh", "--channel", "LTS"); err != nil {
		return fmt.Errorf("dotnet install: %w", err)
	}
	if err := shell.ExecCommand(c, "bash", "-c",
		`echo 'export PATH=$PATH:/config/.dotnet' | sudo -E tee -a /etc/bash.bashrc > /dev/null`,
	); err != nil {
		return fmt.Errorf("dotnet PATH: %w", err)
	}

	logger.Info("Cleaning up...")
	shell.ExecSudoCommand(c, true, "apt-get", "clean")
	shell.ExecSudoCommand(c, true, "rm", "-rf", "/tmp/*", "/var/lib/apt/lists/*", "/var/tmp/*")

	logger.Info("dotnet done.")
	return nil
}

// GoLang installs the given Go version, or the latest if version is empty.
func GoLang(c context.Context, version string) error {
	if version == "" {
		logger.Info("Searching for latest go release...")
		v, err := latestGoVersion()
		if err != nil {
			return fmt.Errorf("golang: could not determine latest version: %w", err)
		}
		version = v
	}

	logger.Info(fmt.Sprintf("Downloading golang (%s)...", version))
	url := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
	if err := shell.ExecCommand(c, "curl", "-#fSL", url, "-o", "/tmp/golang.tar.gz"); err != nil {
		return fmt.Errorf("golang download: %w", err)
	}

	logger.Info(fmt.Sprintf("Installing golang (%s)...", version))
	if err := shell.ExecSudoCommand(c, false, "rm", "-rf", "/usr/local/go"); err != nil {
		return fmt.Errorf("golang rm: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "tar", "-C", "/usr/local", "-xzf", "/tmp/golang.tar.gz"); err != nil {
		return fmt.Errorf("golang extract: %w", err)
	}
	if err := shell.ExecCommand(c, "bash", "-c",
		`echo 'export PATH=$PATH:/usr/local/go/bin' | sudo -E tee -a /etc/bash.bashrc > /dev/null`,
	); err != nil {
		return fmt.Errorf("golang PATH: %w", err)
	}

	shell.ExecCommand(c, "rm", "-f", "/tmp/golang.tar.gz")
	logger.Info("done.")
	logger.Info("Please reload bash profile to finalize.")
	return nil
}

// Bun installs the latest Bun version.
func Bun(c context.Context) error {
	logger.Info("Installing Bun requirements...")
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("bun apt-get update: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "--no-install-recommends", "-y", "unzip"); err != nil {
		return fmt.Errorf("bun unzip: %w", err)
	}
	shell.ExecSudoCommand(c, true, "apt-get", "clean")

	logger.Info("Installing Bun binaries...")
	if err := shell.ExecCommand(c, "bash", "-c", "curl -#fSL https://bun.sh/install | bash"); err != nil {
		return fmt.Errorf("bun install: %w", err)
	}

	logger.Info("Adding bun binary to profile...")
	shell.ExecCommand(c, "bash", "-c",
		`echo 'export BUN_INSTALL=$HOME/.bun' | sudo -E tee -a /etc/bash.bashrc > /dev/null`)
	shell.ExecCommand(c, "bash", "-c",
		`echo 'export PATH=$BUN_INSTALL/bin:$PATH' | sudo -E tee -a /etc/bash.bashrc > /dev/null`)

	logger.Info("done.")
	logger.Info("Please reload bash profile to finalize.")
	return nil
}

// Mise installs the mise version manager, activates it in the system bash
// profile, and registers its completion script.
func Mise(c context.Context) error {
	logger.Info("Installing mise as version manager...")
	// MISE_INSTALL_HELP=0 silences the installer hint, arkanum wires the
	// activation line into /etc/bash.bashrc itself below.
	if err := shell.ExecCommand(c, "bash", "-c",
		"export MISE_INSTALL_HELP=0 && curl -#fSL https://mise.run | bash",
	); err != nil {
		return fmt.Errorf("mise: %w", err)
	}

	logger.Info("Adding mise to profile...")
	if err := shell.ExecCommand(c, "bash", "-c",
		`echo 'export PATH="$HOME/.local/bin:$PATH"' | sudo -E tee -a /etc/bash.bashrc > /dev/null && `+
			`echo 'eval "$($HOME/.local/bin/mise activate bash)"' | sudo -E tee -a /etc/bash.bashrc > /dev/null`,
	); err != nil {
		return fmt.Errorf("mise activate: %w", err)
	}

	logger.Info("Adding mise completion to profile...")
	// The completion script is self-contained, it does not need the
	// bash-completion framework; sourcing it explicitly mirrors how the
	// Dockerfile registers the arkanum completion. The PATH export above
	// matters here, the script looks up `mise` in $PATH when completing.
	if err := shell.ExecCommand(c, "bash", "-c",
		miseBinary()+` completion bash | sudo -E tee /etc/bash_completion.d/mise > /dev/null && `+
			`echo 'source /etc/bash_completion.d/mise' | sudo -E tee -a /etc/bash.bashrc > /dev/null`,
	); err != nil {
		return fmt.Errorf("mise completion: %w", err)
	}

	logger.Info("done.")
	logger.Info("Please reload bash profile to finalize.")
	return nil
}

// NodeJS installs Node.js LTS via mise (installs mise first if needed).
func NodeJS(c context.Context) error {
	if err := Mise(c); err != nil {
		return err
	}
	// Call mise by absolute path: a fresh install lands in ~/.local/bin which
	// is not yet on the current process' PATH.
	logger.Info("Installing NodeJS LTS via mise...")
	if err := shell.ExecCommand(c, miseBinary(), "use", "--global", "node@lts"); err != nil {
		return fmt.Errorf("nodejs: %w", err)
	}
	logger.Info("done.")
	logger.Info("Please reload bash profile to finalize.")
	return nil
}

// miseBinary returns the path where the official installer places the mise binary.
func miseBinary() string {
	return filepath.Join(shell.HomeDir(), ".local", "bin", "mise")
}

// PowerShell installs the latest PowerShell LTS.
func PowerShell(c context.Context) error {
	logger.Info("Installing PowerShell requirements...")
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("powershell apt-get update: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "--no-install-recommends", "-y",
		"apt-transport-https", "software-properties-common"); err != nil {
		return fmt.Errorf("powershell requirements: %w", err)
	}

	logger.Info("Adding powershell package sources...")
	if err := shell.ExecCommand(c, "bash", "-c",
		`curl -#fSL "https://packages.microsoft.com/config/ubuntu/$(lsb_release -rs)/packages-microsoft-prod.deb" -o /tmp/packages-microsoft-prod.deb`,
	); err != nil {
		return fmt.Errorf("powershell repo download: %w", err)
	}
	if err := shell.ExecSudoCommand(c, false, "dpkg", "-i", "/tmp/packages-microsoft-prod.deb"); err != nil {
		return fmt.Errorf("powershell dpkg: %w", err)
	}
	if err := shell.ExecSudoCommand(c, true, "apt-get", "update"); err != nil {
		return fmt.Errorf("powershell apt-get update 2: %w", err)
	}

	logger.Info("Installing PowerShell...")
	if err := shell.ExecSudoCommand(c, false, "apt-get", "install", "--no-install-recommends", "-y", "powershell-lts"); err != nil {
		return fmt.Errorf("powershell install: %w", err)
	}

	logger.Info("done.")
	return nil
}

// GiteaTools installs the tea CLI and changelog tool.
func GiteaTools(c context.Context) error {
	const (
		teaVersion       = "0.11.1"
		changelogVersion = "main"
	)
	logger.Info("Installing Gitea tools...")

	logger.Info(fmt.Sprintf("Downloading 'changelog' (%s)...", changelogVersion))
	changelogURL := fmt.Sprintf("https://dl.gitea.io/changelog-tool/%s/changelog-%s-linux-amd64", changelogVersion, changelogVersion)
	if err := downloadExecutable(c, changelogURL, "/usr/bin/changelog"); err != nil {
		return fmt.Errorf("gitea changelog: %w", err)
	}
	logger.Info("'changelog' command installed.")

	logger.Info(fmt.Sprintf("Downloading 'tea' (%s)...", teaVersion))
	teaURL := fmt.Sprintf("https://dl.gitea.io/tea/%s/tea-%s-linux-amd64", teaVersion, teaVersion)
	if err := downloadExecutable(c, teaURL, "/usr/bin/tea"); err != nil {
		return fmt.Errorf("gitea tea: %w", err)
	}
	logger.Info("'tea' command installed.")

	logger.Info("done.")
	return nil
}

// LazyGit installs the latest LazyGit binary.
func LazyGit(c context.Context) error {
	logger.Info("Installing latest lazygit version...")
	version, err := latestGithubRelease("jesseduffield/lazygit")
	if err != nil {
		return fmt.Errorf("lazygit version: %w", err)
	}

	url := fmt.Sprintf(
		"https://github.com/jesseduffield/lazygit/releases/latest/download/lazygit_%s_Linux_x86_64.tar.gz",
		version,
	)
	if err := shell.ExecCommand(c, "curl", "-#fSL", url, "-o", "/tmp/lazygit.tar.gz"); err != nil {
		return fmt.Errorf("lazygit download: %w", err)
	}
	if err := shell.ExecCommand(c, "tar", "-C", "/tmp", "-xf", "/tmp/lazygit.tar.gz", "lazygit"); err != nil {
		return fmt.Errorf("lazygit extract: %w", err)
	}

	logger.Info("Installing Lazygit binary...")
	if err := shell.ExecSudoCommand(c, false, "install", "/tmp/lazygit", "/usr/local/bin"); err != nil {
		return fmt.Errorf("lazygit install: %w", err)
	}

	logger.Info("Removing lazygit cache files...")
	shell.ExecCommand(c, "rm", "-f", "/tmp/lazygit.tar.gz", "/tmp/lazygit")

	logger.Info("done.")
	return nil
}

// Just installs the latest just binary from GitHub releases.
func Just(c context.Context) error {
	logger.Info("Installing latest just version...")
	version, err := latestGithubRelease("casey/just")
	if err != nil {
		return fmt.Errorf("just version: %w", err)
	}

	url := fmt.Sprintf(
		"https://github.com/casey/just/releases/download/%s/just-%s-x86_64-unknown-linux-musl.tar.gz",
		version, version,
	)
	if err := shell.ExecCommand(c, "curl", "-#fSL", url, "-o", "/tmp/just.tar.gz"); err != nil {
		return fmt.Errorf("just download: %w", err)
	}
	if err := shell.ExecCommand(c, "tar", "-C", "/tmp", "-xf", "/tmp/just.tar.gz", "just"); err != nil {
		return fmt.Errorf("just extract: %w", err)
	}

	logger.Info("Installing just binary...")
	if err := shell.ExecSudoCommand(c, false, "install", "/tmp/just", "/usr/local/bin"); err != nil {
		return fmt.Errorf("just install: %w", err)
	}

	shell.ExecCommand(c, "rm", "-f", "/tmp/just.tar.gz", "/tmp/just")
	logger.Info("done.")
	return nil
}

// ── helpers ──────────────────────────────────────────────────────────────────

func latestGoVersion() (string, error) {
	resp, err := http.Get("https://go.dev/dl/?mode=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected http status: %s", resp.Status)
	}
	var releases []struct {
		Version string `json:"version"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", err
	}
	if len(releases) == 0 {
		return "", fmt.Errorf("empty release list")
	}
	return strings.TrimPrefix(releases[0].Version, "go"), nil
}

func latestGithubRelease(repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected http status: %s", resp.Status)
	}
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return strings.TrimPrefix(release.TagName, "v"), nil
}

// downloadExecutable fetches url into /tmp, verifies the response, then
// installs it to dest via `sudo install` so root owns the final binary.
func downloadExecutable(c context.Context, url, dest string) error {
	name := filepath.Base(dest)
	tmpPath := filepath.Join("/tmp", name)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status: %s", resp.Status)
	}

	f, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		return fmt.Errorf("create tmp file: %w", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		return fmt.Errorf("write tmp file: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close tmp file: %w", err)
	}

	if err := shell.ExecSudoCommand(c, false, "install", "-m", "0755", tmpPath, dest); err != nil {
		return fmt.Errorf("sudo install to %s: %w", dest, err)
	}
	shell.ExecCommand(c, "rm", "-f", tmpPath)
	return nil
}
