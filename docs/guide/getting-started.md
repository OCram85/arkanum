# Getting Started

## 🚧 Prerequisites

### ⚙️ Container Runtime

You need any host with either

- Docker CE / EE running
- or Docker-CE and configured 'swarm' mode.

### 📦 Get the image

You can download the image from Gitea's embedded container registry: `gitea.ocram85.com/arkanum/arkanum` with these tags:

- `latest` - Is based on the lasted master branch commit.
- `next` - Is a test build based on the pull request
- `1`, `0.1`, `0.1.0`, `1.0.0` - tag based version.

> **💡 NOTE: See the [packages page](https://gitea.ocram85.com/arkanum/-/packages/container/arkanum/latest) for latest version and all other available tags.**

The container images are also published to these registries:

- [Docker Hub](https://hub.docker.com/r/ocram85/arkanum)
  - Pull Endpoint: `ocram85/arkanum`
- [GitHub Container Registry](https://github.com/OCram85/arkanum/pkgs/container/arkanum)
  - Pull Endpoint: `ghcr.io/ocram85/arkanum`
- [Codeberg Packages](https://codeberg.org/codeserver/-/packages/container/arkanum/next)
- Pull Endpoint: `codeberg.org/codeserver/arkanum`

## 🏗️ Installation

### Run as Docker Swarm Stack

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
    image: gitea.ocram85.com/arkanum/arkanum:1.0.0
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

### Use Docker-Compose

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

## 🦶 First Steps

After summon Arkanum your first steps should be to set your username and email in the git config:

```bash
arkanum git setup "my-name" "my-email"
```

And that's it. Now you're ready use arkanum as your daily remote code editor. 😄

## 🖼️ Screenshots

![screen1](/screens/screen1.png 'Arkanum startpage')
