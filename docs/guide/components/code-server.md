# Code Server Customizations

## About

The Arkanum image itself depends on the underlying [coder/code-server](https://github.com/coder/code-server) project
and the container image provided by
[linxuserver/docker-code-server](https://github.com/linuxserver/docker-code-server). To be able to add features and
change the behaviour we need to customize either the code-server itself of the docker image creation.

These are the custimzations we added to build the `arkanum server` and the `arkanum-cli` :

## Bash Aliases

You can use the alias `summon` to interact with the code-server instance.

## Product Images

We replaces the product image files to math the arkanum branding. See the
[code-server/media](https://gitea.ocram85.com/arkanum/arkanum/src/branch/master/code-server/media) folder for the
used files.

## Application Name

The code-server command provides the ability to change it's name in the title other places.
We also use arkanum therefore.

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#code-server{Dockerfile:line-numbers}
<<< @/../code-server/root/etc/s6-overlay/s6-rc.d/svc-code-server/run{bash:line-numbers}
:::
