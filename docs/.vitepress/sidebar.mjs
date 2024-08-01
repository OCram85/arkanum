export const SidebarItems = {
  '/guide/': { base: '/guide/', items: getGuide() },
}

function getGuide() {
  return [
    {
      text: 'Guide',
      //collapsed: false,
      items: [
        { text: 'About', link: 'about' },
        { text: 'Getting Started', link: 'getting-started' },
        {
          text: 'Components',
          items: [
            { text: 'Base Images', link: 'components/base-images' },
            { text: 'Included Packages', link: 'components/packages' },
            { text: 'Starship Prompt', link: 'components/starship' },
            { text: 'Git', link: 'components/git' },
            { text: 'Arkanum CLI', link: 'components/arkanum-cli' },
            { text: 'FiraCode Font', link: 'components/firacode' },
            { text: 'Code Server', link: 'components/code-server' },
            { text: 'VSCode', link: 'components/vscode' },
          ],
        },
        { text: 'Known Issues', link: 'known-issues' },
        { text: 'FAQs', link: 'faq' },
      ],
    },
  ]
}
