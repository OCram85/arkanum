---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "Arkanum"
  text: 'An opinionated VS Code distribution'
  tagline: 'Start coding remotely with a containerized editor.'
  image:
    src: /logo.png
    alt: Arkanum
  actions:
    - theme: brand
      text: 🤖 Get Started
      link: /guide/about
    - theme: alt
      text: View on Gitea
      link: https://gitea.ocram85.com/CodeServer/arkanum

features:
  - title: Beginner Friendly
    icon: ✨
    details: Already contains basic extensions and theme setup for quick start.
  - title: Access Anywhere
    icon: 🌎
    details: Access your dev environment from anywhere.
  - title: arkanum cli
    icon: 🧙
    details: Includes <code>arkanum cli</code> to install additional languages and frameworks.
  - title: Gitea + CI/CD
    icon: 🚧
    details: Optimized tooling for Gitea + Woodpecker-CI workflows.
---
