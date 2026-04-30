---
aside: false
author: OCram85
title: 'Arkanum 1.12.0 is released'
tag: 'release'
image: 'blogCard.png'
date: '2026-04-30'
#featured: true
---

# Arkanum 1.12.0 is released

<BlogHeaderImage Source='/blogCard.png' />

**Content**

[[TOC]]

## About

I'm happy to present the latest Arkanum release version `1.12.0`.

This is a substantial update - the bundled code-server jumps from `v4.108.2` all the way to `v4.117.0`,
pulling in nine VS Code releases worth of improvements. On top of that, Arkanum now ships with a
focused, AI-free default configuration and a smarter Go installer.

## Bundled VS Code: 1.108 -> 1.117

The code-server upgrades in this release span **VS Code 1.108 (December 2025) through 1.117 (April 2026)**,
covering a lot of ground. Here are the most notable editor improvements that landed during that window:

**Weekly releases** - Starting with VS Code 1.111, Microsoft switched from a monthly to a weekly stable
release cadence. This means faster bug fixes and shorter iteration cycles reach Arkanum users sooner.

**Integrated browser debugging (1.112)** - Developers can now open web apps directly inside VS Code and
start a full debugging session without leaving the editor. Setting breakpoints, stepping through code,
and inspecting variables all work from the integrated browser panel.

**MCP server sandboxing (1.112)** - Locally configured MCP servers can now run in a sandboxed environment
on macOS and Linux, with restricted file system and network access. A nice security improvement for
anyone running local tooling.

**Refreshed default themes (1.113)** - Both the default light and dark themes received a visual update.
If you have been running with the defaults, you will notice a cleaner, more modern look.

**TypeScript 6.0.3 (1.117)** - The latest TypeScript release is now bundled with VS Code. It includes
recovery fixes for import regressions and ships together with the dependency update to TypeScript 6 in
Arkanum's own toolchain (see Changes below).

> [!NOTE] 💬 NOTE
> Across these nine releases, a large share of VS Code's headline features revolved around
> Copilot and AI chat integrations. That is exactly why Arkanum 1.12.0 ships with all AI chat
> features disabled by default - see the next section.

## Disable AI Chat Features

VS Code has been rapidly expanding its built-in AI and agent capabilities across the 1.109-1.117 range:
GitHub Copilot is now built-in and no longer requires a separate extension, agent debug logs landed in
1.116, and BYOK (bring your own key) for Copilot Business and Enterprise arrived in 1.117 - just to name
a few highlights.

Arkanum is designed as a focused, self-hosted coding environment, and all this AI surface area adds
visual noise and unexpected background network activity out of the box. Starting with `1.12.0`, all
AI chat features are **disabled by default** via the settings configuration
([#221](https://gitea.ocram85.com/arkanum/arkanum/pulls/221)).

The environment stays clean, predictable, and fully under your control - no matter what upstream ships.
If you want to opt in to any of these features, you can always re-enable them by adjusting your
`settings.json`.

## Always Install Latest Go Version

The `arkanum install golang` command now automatically fetches and installs the **latest stable Go release**
instead of relying on a hardcoded version ([#220](https://gitea.ocram85.com/arkanum/arkanum/pulls/220)).
This means your Go toolchain stays current without waiting for an Arkanum version bump.

See the [Arkanum CLI docs](https://arkanum.dev/guide/components/arkanum-cli) for details on the
`install golang` command.

## Changes

This release includes the following changes:

### [v1.12.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.12.0) - 2026-04-30

- 🛠️ ENHANCEMENTS
  - extend(settings): disable all AI chat features ([#221](https://gitea.ocram85.com/arkanum/arkanum/pulls/221))
  - extend(cli): install latest golang version per default ([#220](https://gitea.ocram85.com/arkanum/arkanum/pulls/220))
  - extend(umami): update umami tracker id ([#218](https://gitea.ocram85.com/arkanum/arkanum/pulls/218))
- 🤖 DEPENDENCIES
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.117.0 ([#219](https://gitea.ocram85.com/arkanum/arkanum/pulls/219))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.116.0 ([#216](https://gitea.ocram85.com/arkanum/arkanum/pulls/216))
  - Chore(deps): update dependency typescript to v6 ([#217](https://gitea.ocram85.com/arkanum/arkanum/pulls/217))
  - Chore(deps): update dependency rimraf to v6.1.3 ([#214](https://gitea.ocram85.com/arkanum/arkanum/pulls/214))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.109.2 ([#215](https://gitea.ocram85.com/arkanum/arkanum/pulls/215))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.108.2 ([#210](https://gitea.ocram85.com/arkanum/arkanum/pulls/210))
