# Git

## 🔱 About Git Config files

The git configuration is usually done in these 3 different contexts:

- system wide config
  - `/etc/gitconfig`
- global _(per user)_ config
  - `~/.gitconfig`
- local _(per repository)_
  - `<repo-root>/.git/config`

The final config will be merged in this order.

> [!TIP] 💡 TIP
> You can display the final values with `git config --list`

## 📄 Included Config

We already modified the default system config file:

| Key             | Section          | Description                                                                          |
| --------------- | ---------------- | ------------------------------------------------------------------------------------ |
| `editor`        | **[core]**       | Defines code-server as git edit to open files.                                       |
| `autcrlf`       | **[core]**       | Disables automatic line ending conversion. Should be handeled by the editor          |
| `helper`        | **[credential]** | Enables saving plain credentials for git remotes.                                    |
| `filesEncoding` | **[i18n]**       | Sets _utf-8_ as default encoding.                                                    |
| `default`       | **[push]**       | Uses _simple_ branch name matching strategy.                                         |
| `lg1`           | **[alias]**      | Adds alternate log output layout in graph style                                      |
| `lg2`           | **[alias]**      | Adds extended graph log layout.                                                      |
| `cfetch`        | **[alias]**      | Adds alias for fetch with _--prune_ and _--tags_ . _cfetch_ stands for _clean fetch_ |

There are also some [Phabricator](https://en.wikipedia.org/wiki/Phabricator) workflow inspired helpers:

| Key                | Section     | Description                                                                    |
| ------------------ | ----------- | ------------------------------------------------------------------------------ |
| `feature <branch>` | **[alias]** | Starts a new feature branch from updates master and checks out the new branch. |
| `wip`              | **[alias]** | Stages all current changes and creates a wip _(work in progress)_ commit.      |
| `squish`           | **[alias]** | Takes all open workdir changes and add them to the latest commit.              |
| `pod`              | **[alias]** | Tries to push current branch to _origin/dev_.                                  |
| `poc <branch>`              | **[alias]** | Takes current branch and tries to push it to a new remote one. |

## Git bash Completion

Git bash completion is already enabled in the arkanum image.

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#git{Dockerfile:line-numbers}
<<< @/../gitconfig-system{ini:line-numbers}
:::
