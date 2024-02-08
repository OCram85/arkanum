function getItems(version) {
  const ci = process.env.ci
  const pr = process.env.CI_COMMIT_PULL_REQUEST

  let nver = undefined

  if (ci === 'woodpecker') {
    if (pr !== null) {
      nver = `${version}#${pr}`
    }
  } else {
    nver = version
  }

  return [
    { text: 'Home', link: '/' },
    { text: 'Examples', link: '/markdown-examples' },
    {
      text: nver,
      items: [{ text: 'Changelog', link: '/changelog' }],
    },
  ]
}

export default {
  getItems,
}
