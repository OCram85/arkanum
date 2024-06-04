import { createContentLoader } from 'vitepress'

export default createContentLoader('posts/**/*.md', {
  excerpt: false,
  transform(raw) {
    return raw.sort((a, b) => {
      return +new Date(b.frontmatter.date) - +new Date(a.frontmatter.date)
    })
  }
})
