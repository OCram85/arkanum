## [v1.5.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.5.0) - 2024-02-09

- 🛠️ ENHANCEMENTS
  - Always install latest extension version (#85)
- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.20.1 (#88)
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.20.0 (#86)
- ⚙️ META
  - Add more gitignore items (#89)

## [v1.0.2](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.0.2) - 2023-12-18

- ✨ FEATURES
  - Add bun installer (#81)
- 🛠️ ENHANCEMENTS
  - Bump golang version (#80)
  - Bump default extensions version (#76)
- 📦 BUILD
  - Upd/woodpeckerPluging (#79)
- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.19.1 (#77)
- 📚 DOCS
  - Fix typo in arkanum cli help (#78)
- ⚙️ META
  - Updatest gitea meta files (#75)

## [v1.0.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.0.1) - 2023-11-12

- 🐛 BUGFIXES
  - Fix pwsh install package (#72)
- 📦 BUILD
  - Avoid duplicate ci runs (#73)
- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.18.0 (#70)

## [v1.0.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.0.0) - 2023-09-15

- 🐛 BUGFIXES
  - Fix git config arguments (#68)
  - Fix NodeJs install bug (#67)
  - Fix pwsh install when called as first command (#60)
  - Fix missing volta command (#58)
- 🛠️ ENHANCEMENTS
  - Extends arkanum command and option structure (#62)
- 📦 BUILD
  - Add addiontal container image deployment targets (#63)
- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.16.1 (#59)
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.15.0 (#57)

## [v0.4.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.4.0) - 2023-07-19

- 🐛 BUGFIXES
  - Fix shellcheck issues (#55)
  - Disable starship python module (#52)
- ✨ FEATURES
  - Add docker-cli (#54)
  - Adds Volta as default version manager for NodeJs (#53)
- 🛠️ ENHANCEMENTS
  - Bump go version (#49)

## [v0.3.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.3.1) - 2023-07-04

- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.14.1 (#47)

## [v0.3.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.3.0) - 2023-03-30

- 🛠️ ENHANCEMENTS
  - Add tea cli in gitea block (#41)
- 📦 BUILD
  - Remove repo defined renovate (#44)
- 🤖 DEPENDENCIES
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.11.0 (#45)
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.10.0 (#43)

## [v0.2.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.2.0) - 2023-01-04

- 🐛 BUGFIXES
  - fix typo (#39)
- 🛠️ ENHANCEMENTS
  - Add proxy support (#38)
- 🤖 DEPENDENCIES
  - Update renovate/renovate Docker tag to v34.82 (#37)
  - Update quay.io/linuxserver.io/code-server Docker tag to v4.9.1 (#31)

## [v0.1.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.1.1) - 2023-01-02

- 📦 BUILD
  - fix woodpecker fileMatch regex (#30)
- 🤖 DEPENDENCIES
  - Update renovate/renovate Docker tag to v34.77 (#35)
  - fix renovate config keys (#34)
  - Update renovate/renovate Docker tag to v34.62 (#33)
  - Update renovate/renovate Docker tag to v34.57 (#32)

## [v0.1.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.1.0) - 2022-11-25

- 📦 BUILD
  - sync pipeline build args (#28)
- 📚 DOCS
  - adds Readme content to prepare Github mirror (#26)

## [v0.0.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v0.0.1) - 2022-11-22

- 🐛 BUGFIXES
  - fixes logo size (#22)
- ✨ FEATURES
  - adds FiraCode NerdFont (#9)
  - adds setCode helper (#6)
  - add system gitconfig (#4)
  - add bash-completion and motd handling (#2)
- 🛠️ ENHANCEMENTS
  - Splits extension install into seperate function (#23)
  - Adds readme content (#21)
  - Rename install script to arkanum (#18)
  - adds missing packages required by dotnet (#3)
- 📦 BUILD
  - set image labels (#24)
- 🤖 DEPENDENCIES
  - Adds renovate-bot (#11)
  - update baseimage 4.8.3 (#5)
- 📚 DOCS
  - update refs in Readme (#17)
  - Updates Readme content (#7)
- ⚙️ META
  - adds AGPLv3 license (#20)
  - Adds gitea changelog config (#19)
  - rename project to Arkanum (#16)
  - update PR template wording (#15)
  - Add woodpecker manager in renovate-bot (#14)
  - fixes logo file path (#8)
  - Adds basic container setup with pipeline (#1)
