# Arkanum-cli

## About

The `arkanum` container image includes the `arkanum-cli` to manage the current instance and be able to install
additional resources. To keep the base image as small as possible we decided not to include all possible frameworks.
With this approach the user can decide which frameworks are needed and install it by yourself. This avoids growing
the container image size.

For now the arkanum-cli is made with a simple bash script. But this could change whenever the complexity increases
significantly.

## Usage

To use the cli just open a terminal and call the `arkanum` command. We also added a bash-completion
script for the cli.

<<< @/../arkanum#usage{19-22 bash:line-numbers}

The first thing you normally after starting your Arkanum instance is setting the git user and email:

`arkanum git setup "foobar" "foobar@arkanum.dev"`

Now you can simple add your needed tools and frameworks:

```bash
arkanum install golang
arkanum install nodejs
```

### Session Handling <Badge type="tip" text="^1.8.0" />

The session helper simplifies restoring your session after your instance was restarted. You can now save a global
session config which contains your needed frameworks installed by `arkanum install`. This wraps multiple install
calls and restores your session by running `arkanum session restore`.

For Example:

```bash
# Create a new session config by adding items
arkanum session save lazygit gitea golang
# Shows your saved session
arkanum session show
# Restore your session
arkanum session restore
# optionally reset / delete your session config
arkanum session reset
```

> [!WARNING]
> If your run `arkanum session save` multiple times and use the same arguments, they will also be added and
> installed multiple times!

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#cli{Dockerfile}

<<< @/../arkanum{bash}

<<< @/../arkanum-completion{bash}
:::
