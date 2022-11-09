<p align="right">
  <img src="http://forthebadge.com/images/badges/built-with-love.svg">
  <img src="http://forthebadge.com/images/badges/for-you.svg">
</p>

<p align="center">
  <a href="https://gitea.ocram85.com/CodeServer/Container/">
    <img
      src="/CodeServer/Container/raw/branch/master/assets/social-logo.png"
      alt="Container"
    >
  </a>
</p>

<h1 align="center">
  Code-Server as Container 📦
</h1>

<p align="center">
 This project optimizes the Code-Server container image from linuxserver.io.
</p>

<p align="center">
  <a href="https://ci.ocram85.com/CodeServer/Container">
    <img src="https://ci.ocram85.com/api/badges/CodeServer/Container/status.svg" alt="Master Branch Build Status">
  </a>
</p>

## :book: General

The container is based on the latest `linuxserver/code-server` image.

### 🚀 Starship prompt

The [Starship](starship.rs) prompt is added an enabled as default. Default config uses Emojis and FiraCode icons.

### 🔱 git config

Adds default system config with:

- enabled plain credential store for remote.
- enabled bash completion for git command.
- added git log helper `lg1` + `lg2`

> 💡 See [gitconfig-system](./gitconfig-system) for details.

### 🧙 Added `install-devruntime` helper script

Added `install-devruntime` to help installing common runtime in container.
This helps reducing the image size.

```
install-devruntime is used to install optional runtimes for developing in a
code-server container environment.

Syntax: install-devruntime RUNTIME ...
RUNTIME         [dotnet|golang|nodejs|powershell]
  dotnet        Installs latest LTS dotnet core sdk + runtime.
  golang        Installs golang 1.19.3.
  nodejs        Installs latest NodeJs LTS version.
  powershell    Installs latest PowerShell LTS version.
-h              Prints this help messagee.
disablemotd     Disables hint in new bash terminal
setcode         Sets VS Code user setting with basic (Fira Code)

Example 1: install-devruntime dotnet
Example 2: install devruntime golang nodejs
```

## 💳 Credits

- Photo by <a href="https://unsplash.com/@frankiefoto?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">frank mckenna</a> on <a href="https://unsplash.com/s/photos/container?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>

