export const SidebarItems = {
  '/guide/': { base: '/guide/', items: getGuide() },
}

function getGuide() {
  return [
    {
      text: 'Guide',
      items: [
        { text: 'About', link: '/about' },
        { text: 'Getting Started', link: '/getting-started#Quickstart' },
      ],
    },
  ]
}
