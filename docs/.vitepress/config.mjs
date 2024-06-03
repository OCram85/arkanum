import { defineConfig } from 'vitepress'

import { version } from '../../package.json'
import navbar from './navbar.mjs'
import { SidebarItems } from './sidebar.mjs'
import { brandIcons } from './theme/icons.mjs'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Arkanum',
  titleTemplate: '🧙 Arkanum',
  description: 'An opinionated Code-Server distribution.',
  lang: 'en-US',
  head: [
    ['link', { rel: 'icon', href: '/favicon.png' }],
    [
      'script',
      {
        async: '',
        defer: '',
        src: 'https://umami.ocram85.com/script.js',
        'data-website-id': 'e76d99a0-a185-4151-8a68-0fb44c6f0dc0',
        'data-domains': 'arkanum.dev',
        'data-do-not-track': 'true',
      },
    ],
  ],
  cleanUrls: true,
  appearance: 'force-dark',
  markdown: {
    image: {
      lazyLoading: true,
    },
  },
  themeConfig: {
    logo: '/logo.png',
    outline: 'deep',
    search: {
      provider: 'local',
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: navbar.getItems(version),
    sidebar: SidebarItems,
    socialLinks: [{ icon: {svg: brandIcons.giteaAlt}, link: 'https://gitea.ocram85.com/arkanum/arkanum' }],
    footer: {
      message: 'Released under the AGPLv3 License.',
      copyright: 'Copyright © 2022-present OCram85 <me@ocram85.com>',
    },
    editLink: {
      pattern: 'https://gitea.ocram85.com/arkanum/arkanum/_edit/main/docs/:path',
      text: 'Edit this page on Gitea',
    },
  },
})
