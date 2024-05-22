# :rocket: Starship Prompt

## About

Starship prompt is a powerful addon to modify the shell prompt dynamically based on the current location and state.
Therefore we already install the latest version while building the Arkanum image.
Additionally we use starship as default prompt in bash.

The included default setup uses the system wide installed [NerdFont](https://www.nerdfonts.com/) patched [FiraCode](https://github.com/tonsky/FiraCode) Font with its symbols, ligatures and emojis.

## Screenshots

![screen1](./prompts/prompt1.png 'gtit repo with nodejs + bun symbols and package.json version displayed')

![screen2](./prompts/prompt2.png 'default prompt outside of a git repo')

![screen3](./prompts/prompt3.png 'local git repo with two commits behind origin')

![screen4](./prompts/prompt4.png 'same repo prompt after pulling')

![screen5](./prompts/prompt5.png 'git rep with golang project and open workspace changes')

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#starship{Dockerfile:line-numbers}

<<< @/../starship.toml{toml:line-numbers}
:::
