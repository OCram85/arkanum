# Base Images

## Coder/code-server

Arkanum is build on top of the [code-server](https://github.com/coder/code-server) project. This project
provides the base ability to run VS Code remotely and accessible by any browser instead of the need for a local
installation.

The folks from [linuxserver.io](https://www.linuxserver.io/) build and maintainer super reliable container images
for all kind of software. And so they also did for the code-server project.

We also took their [code-server image](https://hub.docker.com/r/linuxserver/code-server) as our base for further
customizing and tailoring arkanum.

## gitpod-io/openvscode-server <Badge type="info" text="planned for arkanum v2++" />

> [!NOTE] 💬 NOTE
> We plan to also build Arkanum with the openvscode-server base. This could enable us to install Arkanum as
> PWA with our own branding like the product logo.
