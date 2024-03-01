# What is Arkanum

### 🚀 Starship prompt

We added the [Starship](./components/starship.md) prompt is as default in the integrated terminal. The default config uses Emojis and FiraCode icons.

### 🔱 git config

Added default git system config file with:

- code-server as default editor.
- disabled `aurocrlf`.
- enabled plain credential store for remote.
- added git log helper `lg1` + `lg2`.
- enabled bash completion for git command in integrated bash terminal.

> 💡 See [gitconfig-system](./components/git.md) for details.

### 🧙 `arkanum` helper

Added `arkanum` to help installing common runtimes in container.
This helps reducing the image size.

```
🧙 arkanum ✨🌌☄️💥 is used to install optional tools for developing in a
  code-server container environment.

  Syntax: arkanum <flags> COMMAND OPTION ARGUMENT
  COMMAND
    config                 The config command is used to modify arkanum itself.
    git                    The git command is a wrapper for git helpers.
    install                The install command is used to add different tools
    help                   Shows this help text.

  OPTION
    config:
      disable-motd         Disables hint in new bash terminal.
      install extensions   Installs predefined recommended extensions.
      reset-codesettings   Sets VS Code user setting with basic (Fira Code).

    git:
      setup                Takes two arguments to setup the git client:
                             1) Username
                             2) Email address

    install:
      docker-cli           Installs the latest docker-cli.
      dotnet               Installs latest LTS dotnet core sdk + runtime.
      gitea                Installs gitea tools like the changelog generator.
      golang               Installs golang 1.19.3.
      bun                  Installs latest bun version.
      nodejs               Installs latest NodeJs LTS version using Volta.
      volta                Installs Volta as NodeJS version manager.
      powershell           Installs latest PowerShell LTS version.

  Example 1: arkanum git setup "my-name" "my-email"
  Example 2: arkanum install golang
  Example 3: arkanum config disable-motd
```

### 📝 Fira Code (NerdFont patched)

Added FiraCode as default font in editor and integrated terminal. The font files are embedded and can be used without local installation.

### 🦸 VSCode default settings

If your start the container or log in the first time, a default config file is deployed.

This user setting defines the following stuff:

- Use compact menu bar to avoid users with multiple menu bars.
- Use _One Dark Pro Darker_ theme
- Use _vscode-icons_ icon set
- Set FiraCode as default font in editor.
  - Tries to use alternate font names for FiraCode if its locally available.
- Sets FiraCode mono variant in terminal to enable icons used by starship prompt.
- Enables font ligatures
- Enables _auto save_ and _format on save_.
- Disables auto update for extension.
- Disables VSCode telemetry
- Disable confirm message for sync branches.

Additionally we install these extensions on container startup:

- [One Dark Pro](https://open-vsx.org/extension/zhuangtongfa/material-theme) theme
- [vscode-icons](https://open-vsx.org/extension/vscode-icons-team/vscode-icons) icon set
- [Gitlens](https://open-vsx.org/extension/eamodio/gitlens)

## 💣 Known Issues

### Starship.rs

Starship detects workspaces as active python projects. It always appends the prompt
fragment `via 🐍 (lsiopy)`. For now I disabled the python module in starship.

### 🐛 Default extensions installation timing error

If the automatic installation of the default extension fails, you can always retry he installation with the
following command:

```bash
# restart the installation
arkanum config install-extensions
# Optional: reset the vscode user setting
arkanum config reset-codesettings
# Reload with command F1 + Developer: Reload Window
```
