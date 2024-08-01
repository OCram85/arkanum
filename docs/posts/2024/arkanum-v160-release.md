---
aside: false
author: OCram85
title: 'Arkanum 1.6.0 is released'
tag: 'release'
image: 'blogCard.png'
date: '2024-08-01'
#featured: true
---

<!-- markdownlint-disable MD025 MD033 MD036 -->

# Arkanum 1.6.0 is released

<BlogHeaderImage Source='/blogCard.png' />

**Content**

[[TOC]]

## About

I'm happy to present the latest Arkanum release version `1.6.0`.

This release includes the following changes:

## Product Image Update

We updated the Arkanum images to be able to generate the favicons and PWA icons for the code-server image.

These are the curren image iterations:

|                                          `pre v1`                                           |                                        `v1++`                                         |      `v1.6++` + `v2++`     |
| :-----------------------------------------------------------------------------------------: | :-----------------------------------------------------------------------------------: | :----------------: |
| ![pre-v1](https://gitea.ocram85.com/arkanum/arkanum/raw/branch/master/assets/logo-pre1.png) | ![v1](https://gitea.ocram85.com/arkanum/arkanum/raw/branch/master/assets/logo-v1.png) | ![v1++](/logo.png) |

> [!TIP] 💡 TIP
> If you're interested in all variations, you can find the sources at
> Canva -> [Arkanum 2.0 Collection](https://www.canva.com/design/DAGMBuM5uTk/mIyXxRbPwS6ZiT7I-MVepQ/view?utm_content=DAGMBuM5uTk&utm_campaign=designshare&utm_medium=link&utm_source=editor)

## Starting a Matrix Space

We decided to use Matrix for building the Arkanum community.

Feel free to join the [public community space](https://matrix.to/#/#arkanum:matrix.org).

[![Chat on Matrix](https://matrix.to/img/matrix-badge.svg)](https://matrix.to/#/#arkanum:matrix.org)

We also think about hosting a _Discourse_ or _Apache Answer_ based forum whenever needed. This could help collecting
all useful community resources like discussions, FAQs, and guides without searching in a room history.

## Added Tools <Badge type="tip" text="v1.6.0++" />

- We added [lazygit](https://github.com/jesseduffield/lazygit) in arkanum-cli.
- We also added the bash alias `summon` for _code-server_.

## Dependencies

- Updated Code-Server to `v4.91.0`
- Updated dependencies used in WoodpeckerCI workflows.

## Changelog

<!--@include: ../../../CHANGELOG.md{3,15}-->
