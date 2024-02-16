import { defineConfig } from 'vitepress'

import { devDependencies, version } from '../../package.json'
import navbar from './navbar.mjs'
import { SidebarItems } from './sidebar.mjs'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Arkanum',
  titleTemplate: '🧙 Arkanum',
  description: 'An opinionated Code-Server distribution.',
  lang: 'en-US',
  head: [['link', { rel: 'icon', href: '/favicon.png' }]],
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
      provider: 'local'
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: navbar.getItems(version),
    sidebar: SidebarItems,
    socialLinks: [{ icon: 'github', link: 'https://gitea.ocram85.com/CodeServer/arkanum' }],
    footer: {
      message: 'Released under the AGPLv3 License.',
      copyright: 'Copyright © 2022-present OCram85 <me@ocram85.com>'
    },
    editLink: {
      pattern: 'https://gitea.ocram85.com/CodeServer/arkanum/_edit/main/docs/:path',
      text: 'Edit this page on Gitea'
    }
  },
})
