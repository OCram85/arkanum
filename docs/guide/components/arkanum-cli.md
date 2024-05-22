# Arkanum-cli

## About

The `arkanum` container image includes the `arkanum-cli` to manage the current instance and be able to install
additional resources. To keep the base image as small as possible we decided not ot include all possible frameworks. This decision also helps to give the use the ability to manage the needed framework versions like for
NodeJS or Golang.

For now the arkanum-cli is made with a simple bash script. But this could change whenever the complexity increases
significantly.

## Usage

To use the cli just open a terminal and call the `arkanum` command. We also added a bash-completion
script for the cli.

<<< @/../arkanum#usage{34-36 bash:line-numbers}

The first thing you normally after starting your Arkanum instance is setting the git user and email:

`arkanum git setup "foobar" "foobar@arkanum.dev"`

Now you can simple add your needed tools and frameworks:

```bash
arkanum install golang
arkanum install nodejs
```

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#cli{Dockerfile}

<<< @/../arkanum{bash}

<<< @/../arkanum-completion{bash}
:::
