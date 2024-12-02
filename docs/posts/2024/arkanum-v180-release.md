---
aside: false
author: OCram85
title: 'Arkanum 1.8.1 is released'
tag: 'release'
image: 'blogCard.png'
date: '2024-12-02'
#featured: true
---

# Arkanum 1.8.1 is released

<BlogHeaderImage Source='/blogCard.png' />

**Content**

[[TOC]]

## About

I'm happy to present the latest Arkanum release version `1.8.1`.

We highly encourage you to directly update to version `1.8.1`. This hotfix version already
fixes an issue where the embedded FiraCode font couldn't be loaded.

## New Session Handler

This version brings you a basic session handler. It includes saving your preferred framework
sources and simplifies the restore after an Arkanum restart.

See the docs for details [here](../../guide/components/arkanum-cli#session-handling).

## Changes

This release includes the following changes:

### [v1.8.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.8.1) - 2024-12-02

- 🐛 BUGFIXES
  - Fix FiraCode font loader (#147)
- 🤖 DEPENDENCIES
  - Chore(deps): update dependency prettier to v3.4.1 (#148)
  - Chore(deps): update dependency typescript to v5.7.2 (#145)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.95.3 (#144)
  - Chore(deps): update oven/bun docker tag to v1.1.38 (#143)

### [v1.8.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.8.0) - 2024-11-21

- ✨ FEATURES
  - Add session helper (#141)
- 🛠️ ENHANCEMENTS
  - Move additional packages install to dockerfile (#138)
- 🤖 DEPENDENCIES
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.95.2 (#139)
  - Chore(deps): update dependency vitepress to v1.5.0 (#137)
- ⚙️ META
  - Pin exact bun version (#140)
