## [v1.8.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.8.0) - 2024-11-21

- ✨ FEATURES
  - Add session helper (#141)
- 🛠️ ENHANCEMENTS
  - Move additional packages install to dockerfile (#138)
- 🤖 DEPENDENCIES
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.95.2 (#139)
  - Chore(deps): update dependency vitepress to v1.5.0 (#137)
- ⚙️ META
  - Pin exact bun version (#140)

## [v1.7.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.7.1) - 2024-10-23

- 🐛 BUGFIXES
  - Move file binary install to powershell command (#134)
- 🛠️ ENHANCEMENTS
  - Bump golang version (#135)

## [v1.7.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.7.0) - 2024-10-22

- 🐛 BUGFIXES
  - Adds file binary to supplement fix Publish-Module pwsh command. (#132)
- 🤖 DEPENDENCIES
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v5 (#131)
  - Chore(deps): update dependency typescript to v5.6.3 (#129)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.93.1 (#127)
  - Chore(deps): update dependency typescript to v5.6.2 (#126)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.92.2 (#125)
  - Chore(deps): pin dependencies (#124)

## [v1.6.0](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.6.0) - 2024-07-31

- ✨ FEATURES
  - Add arkanum branding for code-server (#120)
  - Add lazygit (#116)
- 🛠️ ENHANCEMENTS
  - Update brand assets (#121)
  - Migrate codeberg organization to arkanum (#117)
- 🤖 DEPENDENCIES
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v4.2.0 (#119)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.91.1 (#118)
  - Chore(deps): update dependency rimraf to v6 (#115)
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v4.1.0 (#114)
  - Add renovate reviewer (#113)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.91.0 (#112)

## [v1.5.3](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.5.3) - 2024-06-20

- 🛠️ ENHANCEMENTS
  - Update blog + bump golang (#108)
- 🤖 DEPENDENCIES
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v4 (#110)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.90.2 (#109)
  - Chore(deps): update caddy docker tag to v2.8.4 (#106)
- 📚 DOCS
  - Extend blog card layout (#107)

## [v1.5.2](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.5.2) - 2024-06-04

- 🐛 BUGFIXES
  - Fix typo (#104)
- ✨ FEATURES
  - Add blog feature to docs (#101)
- 🛠️ ENHANCEMENTS
  - Add gitea icon (#102)
- 🤖 DEPENDENCIES
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.89.1 (#100)

## [v1.5.1](https://gitea.ocram85.com/arkanum/arkanum/releases/tag/v1.5.1) - 2024-05-22

- 🐛 BUGFIXES
  - Fix starship prompt installation (#95)
- ✨ FEATURES
  - Add Arkanum.dev site with docs (#97)
- 🤖 DEPENDENCIES
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v4 (#98)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.89.0 (#96)
  - Chore(deps): update woodpeckerci/plugin-docker-buildx docker tag to v3 (#93)
  - Chore(deps): update quay.io/linuxserver.io/code-server docker tag to v4.22.1 (#91)

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
