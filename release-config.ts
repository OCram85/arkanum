export default {
  skipLabels: [
    // FIXME: ignoring a label prevents processing the full PR!
    //'build_pr_images',
    // ignore issue labels
    'duplicate',
    'help wanted',
    'invalid',
    'question',
    'wontfix',
  ],
  changeTypes: [
    {
      title: '💥 Breaking changes',
      labels: ['breaking'],
      bump: 'major',
      weight: 4,
    },
    {
      title: '✨ Features',
      labels: ['feature'],
      bump: 'minor',
      weight: 3,
    },
    {
      title: '🛠️ Enhancement',
      labels: ['enhancement'],
      bump: 'minor',
      weight: 2,
    },
    {
      title: '🐛 Bug Fixes',
      labels: ['bug'],
      bump: 'patch',
      weight: 1,
    },
    {
      title: '🤖 Dependency',
      labels: ['renovate', 'dependency'],
      bump: 'patch',
    },
    {
      title: '📚 Documentation',
      labels: ['docs'],
      bump: 'patch',
    },
    {
      title: '📦 Build',
      labels: ['pipeline'],
      bump: 'patch',
      weight: -1,
    },
    {
      title: '⚙️ Meta',
      labels: ['meta'],
      bump: 'patch',
      weight: -2,
    },
    {
      title: '🔖 Misc',
      labels: ['misc'],
      bump: 'patch',
      default: true,
      weight: -3,
    },
  ],
}
