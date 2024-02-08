import { defineConfig } from 'vitepress'

import { repository, version } from '../../package.json'
import navbar from './navbar.mjs'
import { SidebarItems } from './sidebar.mjs'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Arkanum',
  titleTemplate: '🧙 Arkanum',
  description: 'An opinionated Code-Server distribution.',
  head: [['link', { rel: 'icon', href: '/favicon.png' }]],
  cleanUrls: true,
  appearance: 'dark',
  markdown: {
    image: {
      lazyLoading: true,
    },
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: navbar.getItems(version),
    sidebar: SidebarItems,
    socialLinks: [{ icon: 'github', link: 'https://github.com/vuejs/vitepress' }],
  },
})
