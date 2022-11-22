<p align="right">
  <img src="http://forthebadge.com/images/badges/built-with-love.svg">
  <img src="http://forthebadge.com/images/badges/for-you.svg">
</p>

<p align="center">
  <a href="https://gitea.ocram85.com/CodeServer/arkanum/">
    <img
      src="/CodeServer/arkanum/raw/branch/master/assets/social-logo.png"
      alt="Container"
    >
  </a>
</p>

<h1 align="center">
  🧙 Arkanum ✨ 🌌 ☄️ 💥
</h1>

<p align="center">
Code-Server container optimized for daily use. ❤
</p>

<p align="center">
  <a href="https://ci.ocram85.com/CodeServer/arkanumg">
    <img src="https://ci.ocram85.com/api/badges/CodeServer/arkanum/status.svg" alt="Master Branch Build Status">
  </a>
</p>

## :book: General

The container is based on the latest `linuxserver/code-server` image.

### 🚀 Starship prompt

The [Starship](starship.rs) prompt is added and enabled as default. Default config uses Emojis and FiraCode icons.

### 🔱 git config

Adds default git system config file with:

- code-server as default editor.
- disabled `aurocrlf`.
- enabled plain credential store for remote.
- added git log helper `lg1` + `lg2`.
- enabled bash completion for git command in integrated bash terminal.
> 💡 See [gitconfig-system](./gitconfig-system) for details.

### 🧙 Added `arkanum` helper script

Added `arkanum` to help installing common runtime in container.
This helps reducing the image size.

```
🧙 arkanum ✨🌌☄️💥 is used to install optional runtimes for developing in a
  code-server container environment.

  Syntax: arkanum RUNTIME ...
  RUNTIME         [dotnet|golang|nodejs|powershell]
    dotnet        Installs latest LTS dotnet core sdk + runtime.
    gitea         Installs gitea tools like the changelog generator.
    golang        Installs golang 1.19.3.
    nodejs        Installs latest NodeJs LTS version.
    powershell    Installs latest PowerShell LTS version.
  --disable-motd         Disables hint in new bash terminal.
  --install-extensions   Installs predefined recommended extensions.
  --reset-codesetting    Sets VS Code user setting with basic (Fira Code).
  -h                     Prints this help message.

  Example 1: arkanum dotnet
  Example 2: arkanum golang nodejs
  Example 3: arkanum --disable-motd
```
### 📝 Fira Code (NerdFont patched)

Added FiraCode as default font in editor and integrated terminal. The font files are embedded and can be used without local installation.

### VSCode default settings

If your start the container or log in the first time, a default config file is deployed.

This user setting defines the following stuff:

- Use compact menu bar to avoid users with multiple menu bars.
- Use *One Dark Pro Darker* theme
- Use *vscode-icons* icon set
- Set 'FiraCode' as default font in editor.
  - Tries to use alternate font names for FiraCode if its locally available.
- Sets 'FiraCode' mono variant in terminal to enable icons used by starshop prompt.
- Enables font ligatures
- Enables *auto save* and *format on save*.
- Disables auto update for extension.
- Disables VScode telemetry
- Disable confirm message for sync branches.

Additionally we install these extensions on container startup:

- [One Dark Pro](https://open-vsx.org/extension/zhuangtongfa/material-theme) theme
- [vscode-icons](https://open-vsx.org/extension/vscode-icons-team/vscode-icons) icon set
- [Gitlens](https://open-vsx.org/extension/eamodio/gitlens)

## 💳 Credits

Akranum is based on the following projects and wouldn't be possible without:

- [microsoft/vscode](https://github.com/microsoft/vscode) - Visual Studio Code, OSS. `[MIT]`
- [coder/code-server](https://github.com/coder/code-server) - VSCode on a remote server, accessible through the browser. `[MIT]`
- [linuxserver/docker-code-server](https://github.com/linuxserver/docker-code-server) - docker image based for *coder/code-server*. `[GPL-3.0]`
- A huge thanks to tuanpham for sharing his [code-server font patch](https://github.com/tuanpham-dev/code-server-font-patch).

## ⚖️ License (AGPLv3)

![AGPL](https://www.gnu.org/graphics/agplv3-155x51.png)

```
Arkanum - Code-Server container optimized for daily use.
Copyright (C) 2022  "OCram85 <me@ocram85.com>"

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
