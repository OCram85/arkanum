<p align="center">
  <a href="https://gitea.ocram85.com/arkanum/arkanum/">
    <img
      src="https://gitea.ocram85.com/arkanum/arkanum/raw/branch/master/assets/social-logo.png"
      alt="Container"
    >
  </a>
</p>

<h1 align="center">
  🧙 Arkanum ✨ 🌌 ☄️ 💥
</h1>

<p align="center">
... is a Code-Server container optimized for daily use.
</p>

<p align="center">
  <a href="https://ci.ocram85.com/arkanum/arkanum">
    <img src="https://ci.ocram85.com/api/badges/arkanum/arkanum/status.svg" alt="Master Branch Build Status">
  </a>
</p>

## 🤖 Quickstart

### 1. ⚡ Get the image 📦

You can download the image from the gitea embedded container registry: `gitea.ocram85.com/arkanum/arkanum` with these tags:

- `latest` - Is based on the lasted master branch commit.
- `next` - Is a test build based on the pull request
- `1`, `0.1`, `0.1.0`, `1.0.0` - tag based version.

> **💡 NOTE: See the [packages page](https://gitea.ocram85.com/arkanum/-/packages/container/arkanum/latest) for latest version and all other available tags.**

The container images are also published to these registries:

- [Docker Hub](https://hub.docker.com/r/ocram85/arkanum)
  - Pull Endpoint: `ocram85/arkanum`
- [GitHub Container Registry](https://github.com/OCram85/arkanum/pkgs/container/arkanum)
  - Pull Endpoint: `ghcr.io/ocram85/arkanum`
- [Codeberg Packages](https://codeberg.org/arkanum/-/packages/container/arkanum/next)
- Pull Endpoint: `codeberg.org/arkanum/arkanum`

### 2.a Run as Docker Swarm Stack

This example shows how to run arkanum as an additional swarm stack.

Therefore you need

- an already running docker swarm cluster,
- a running traefik instance handling the http and https routes,
- configured to expose services in the a ingress overlay network called `traefik-public`.

> ❗ **Warning:** Make sure to secure the access to arkanum with proper **authentication method** and use
> a trusted + **secure https connection**.

```yaml
version: '3.8'
services:
  arkanum:
    image: gitea.ocram85.com/arkanum/arkanum:1
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/Berlin
      - PASSWORD=foo #optional
      #- HASHED_PASSWORD= #optional
      - SUDO_PASSWORD=foobar #optional
      #- SUDO_PASSWORD_HASH= #optional
      #- PROXY_DOMAIN=code-server.my.domain #optional
      - DEFAULT_WORKSPACE=/config/workspace
    deploy:
      replicas: 1
      labels:
        - 'traefik.enable=true'
        - 'traefik.docker.network=traefik-public'
        - 'traefik.http.routers.arkanum.rule=Host(`vscode.mydomain.com`)'
        - 'traefik.http.routers.arkanum.tls.certresolver=myresolver'
        - 'traefik.http.services.arkanum-srv.loadbalancer.server.port=8443'
    volumes:
      # store workspace and use config in volume.
      - codedata:/config
      # mount docker socket to manage host docker
      - /var/run/docker.sock:/var/run/docker.sock
    # no need to expose the port. traefik acts as reverse proxy and handles the https access.
    #ports:
    #  - 8443:8443
    networks:
      - arkanum-sphere
      - traefik-public

volumes:
  codedata:

networks:
  arkanum-sphere:
  traefik-public:
    external: true
```

> 💡 NOTE: For advanced config with additional environment variables see [linuxserver/docker-code-server](https://github.com/linuxserver/docker-code-server) help.

### 2.b Use Docker-Compose

This is a basic example for a `docker-compose` file from the [linuxserver/docker-code-server](https://github.com/linuxserver/docker-code-server) project.

See their [docs](https://github.com/linuxserver/docker-code-server#parameters) about a detailed help for advanced config parameters.

```yaml
---
version: '3.8'
services:
  arkanum:
    image: gitea.ocram85.com/arkanum/arkanum:1
    container_name: code-server
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
      - PASSWORD=password #optional
      - HASHED_PASSWORD= #optional
      - SUDO_PASSWORD=password #optional
      - SUDO_PASSWORD_HASH= #optional
      - PROXY_DOMAIN=code-server.my.domain #optional
      - DEFAULT_WORKSPACE=/config/workspace #optional
    volumes:
      - /path/to/appdata/config:/config
    ports:
      - 8443:8443
    restart: unless-stopped
```

### 3. 🦶 First Steps

After summon Arkanum your first steps should be to set your username and email in the git config:

```bash
arkanum git setup "my-name" "my-email"
```

And that's it. Now you're ready use arkanum as your daily remote code editor. 😄

## 📖 Content

<p align="center">
  <a href="https://gitea.ocram85.com/arkanum/arkanum/">
    <img
      src="https://gitea.ocram85.com/arkanum/arkanum/raw/branch/master/assets/screen1.png"
      alt="Screenshot1"
    >
  </a>
</p>

### 🚀 Starship prompt

We added the [Starship](starship.rs) prompt is as default in the integrated terminal. The default config uses Emojis and FiraCode icons.

### 🔱 git config

Added default git system config file with:

- code-server as default editor.
- disabled `aurocrlf`.
- enabled plain credential store for remote.
- added git log helper `lg1` + `lg2`.
- enabled bash completion for git command in integrated bash terminal.

> 💡 See [gitconfig-system](./gitconfig-system) for details.

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
      lazygit              Installs latest Lazygit binary.

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

## 😡 We're Using GitHub Under Protest

This project is currently **mirrored** to GitHub. This is not ideal; GitHub is a
proprietary, trade-secret system that is not Free and Open Source Software
(FOSS). We are deeply concerned about using a proprietary system like GitHub
to develop our FOSS project. We have an
[open Gitea repository](https://gitea.ocram85.com/arkanum/arkanum/issues) where the
project contributors are actively discussing how we can move away from GitHub
in the long term. We urge you to read about the
[Give up GitHub](https://GiveUpGitHub.org) campaign from
[the Software Freedom Conservancy](https://sfconservancy.org) to understand
some of the reasons why GitHub is not a good place to host FOSS projects.

If you are a contributor who personally has already quit using GitHub, please
[check this resource](https://gitea.ocram85.com/arkanum/arkanum) for how to send us contributions without
using GitHub directly.

Any use of this project's code by GitHub Copilot, past or present, is done
without our permission. We do not consent to GitHub's use of this project's
code in Copilot.

![Logo of the GiveUpGitHub campaign](https://sfconservancy.org/img/GiveUpGitHub.png)

## 🙏 Credits

Akranum is based on the following projects and wouldn't be possible without them:

- [microsoft/vscode](https://github.com/microsoft/vscode) - Visual Studio Code, OSS. `[MIT]`
- [coder/code-server](https://github.com/coder/code-server) - VSCode on a remote server, accessible through the browser. `[MIT]`
- [linuxserver/docker-code-server](https://github.com/linuxserver/docker-code-server) - docker image based for _coder/code-server_. `[GPL-3.0]`
- A huge thanks to tuanpham for sharing his [code-server font patch](https://github.com/tuanpham-dev/code-server-font-patch).

## ⚖️ License (AGPLv3)

![AGPL](https://www.gnu.org/graphics/agplv3-155x51.png)

```text
Arkanum - Code-Server container optimized for daily use.
Copyright (C) 2022 "OCram85 <me@ocram85.com>"

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
