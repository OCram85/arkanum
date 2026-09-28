---
aside: false
author: OCram85
title: 'Arkanum 1.13.0 is released'
tag: 'release'
image: 'blogCard.png'
date: '2026-07-31'
#featured: true
---

# Arkanum 1.13.0 is released

<BlogHeaderImage Source='/blogCard.png' />

**Content**

[[TOC]]

## About

I'm happy to present the latest Arkanum release version `1.13.0`.

This update focuses on a significant leap forward for our underlying infrastructure, most notably the bundled
code-server which now jumps from `v4.117.0` all the way to `v4.131.0`. This brings in months of improvements and
stability fixes directly into your environment. Additionally, we've introduced some minor refinements to clean up
the user experience by removing unnecessary prompt hints, continuing our goal of providing a streamlined,
focused workspace.

## Bundled VS Code: 1.117 -> 1.131

The jump in code-server version is substantial. By moving to `v4.131.0`, Arkanum users benefit from the latest
features and optimizations included in that release cycle of the downstream core components.

**Refined Experience** - We have also implemented improvements to remove unnecessary prompt hints (`#234`).
This helps keep your workflow smooth and distraction-free, aligning with our philosophy of a "clean" default setup
where only what you need is presented to you.

> [!NOTE] 💬 NOTE
> Arkanum continues to be designed as a focused, self-hosted coding environment. We aim to provide the stable
> foundation where the tools work reliably in the background without unnecessary noise or unwanted features being
> enabled by default.

## Changes

This release includes the following changes:

### [v1.13.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.13.0) - 2026-07-31

- 🛠️ ENHANCEMENTS
  - extend(settings): disable prompt hint ([#234](https://gitea.ocram85.com/arkanum/arkanum/pulls/234))
- 🤖 DEPENDENCIES
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.131.0 ([#266](https://gitea.ocram85.com/arkanum/arkanum/pulls/266))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.130.0 ([#265](https://gitea.ocram85.com/arkanum/arkanum/pulls/265))
  - Chore(deps): update dependency prettier to v3.9.6 ([#264](https://gitea.ocram85.com/arkanum/arkanum/pulls/264))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.129.0 ([#263](https://gitea.ocram85.com/arkanum/arkanum/pulls/263))
  - Chore(deps): update woodpeckerci/plugin-ready-release-go docker tag to v4.1.2 ([#262](https://gitea.ocram85.com/arkanum/arkanum/pulls/262))
  - Chore(deps): update dependency prettier to v3.9.5 ([#261](https://gitea.ocram85.com/arkanum/arkanum/pulls/261))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.128.0 ([#259](https://gitea.ocram85.com/arkanum/arkanum/pulls/259))
  - Chore(deps): update commitlint/commitlint docker tag to v21.2.1 ([#258](https://gitea.ocram85.com/arkanum/arkanum/pulls/258))
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v6.1.1 ([#256](https://gitea.ocram85.com/arkanum/arkanum/pulls/256))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.127.0 ([#254](https://gitea.ocram85.com/arkanum/arkanum/pulls/254))
  - Chore(deps): update dependency prettier to v3.9.4 ([#255](https://gitea.ocram85.com/arkanum/arkanum/pulls/255))
  - Chore(deps): update dependency prettier to v3.9.3 ([#253](https://gitea.ocram85.com/arkanum/arkanum/pulls/253))
  - Chore(deps): update commitlint monorepo to v21.2.0 ([#251](https://gitea.ocram85.com/arkanum/arkanum/pulls/251))
  - Chore(deps): update dependency prettier to v3.9.1 ([#252](https://gitea.ocram85.com/arkanum/arkanum/pulls/252))
  - Chore(deps): update dependency prettier to v3.8.5 ([#250](https://gitea.ocram85.com/arkanum/arkanum/pulls/250))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.126.0 ([#249](https://gitea.ocram85.com/arkanum/arkanum/pulls/249))
  - Chore(deps): update commitlint monorepo to v21.1.0 ([#248](https://gitea.ocram85.com/arkanum/arkanum/pulls/248))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.125.0 ([#247](https://gitea.ocram85.com/arkanum/arkanum/pulls/247))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.124.2 ([#246](https://gitea.ocram85.com/arkanum/arkanum/pulls/246))
  - Chore(deps): update dependency prettier to v3.8.4 ([#245](https://gitea.ocram85.com/arkanum/arkanum/pulls/245))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.123.0 ([#244](https://gitea.ocram85.com/arkanum/arkanum/pulls/244))
  - Chore(deps): update caddy docker tag to v2.11.4 ([#243](https://gitea.ocram85.com/arkanum/arkanum/pulls/243))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.122.1 ([#242](https://gitea.ocram85.com/arkanum/arkanum/pulls/242))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.122.0 ([#240](https://gitea.ocram85.com/arkanum/arkanum/pulls/240))
  - Chore(deps): update commitlint/commitlint docker tag to v21.0.2 ([#239](https://gitea.ocram85.com/arkanum/arkanum/pulls/239))
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.121.0 ([#224](https://gitea.ocram85.com/arkanum/arkanum/pulls/224))
  - Chore(deps): update dependency prettier to v3.8.3 ([#223](https://gitea.ocram85.com/arkanum/arkanum/pulls/223))
  - Chore(deps): update caddy docker tag to v2.11.3 ([#228](https://gitea.ocram85.com/arkanum/arkanum/pulls/228))
  - Chore(deps): update woodpeckercommunity/peckify docker tag to v0.2.1 ([#227](https://gitea.ocram85.com/arkanum/arkanum/pulls/227))
  - Chore(deps): update oven/bun docker tag to v1.3.14 ([#225](https://gitea.ocram85.com/arkanum/arkanum/pulls/225))
  - Chore(deps): enable automerge ([#235](https://gitea.ocram85.com/arkanum/arkanum/pulls/235))
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v6.1.0 ([#230](https://gitea.ocram85.com/arkanum/arkanum/pulls/230))
  - Chore(deps): update commitlint/commitlint docker tag to v21 ([#231](https://gitea.ocram85.com/arkanum/arkanum/pulls/231))
  - Chore(deps): update woodpeckerci/plugin-ready-release-go docker tag to v4 ([#232](https://gitea.ocram85.com/arkanum/arkanum/pulls/232))
