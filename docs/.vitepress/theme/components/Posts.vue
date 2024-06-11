<script setup>
  import { data } from './posts.data.js'
  import Post from './Post.vue'

  const props = defineProps({
    tagFilter: {
      type: String,
      default: 'none',
    },
    featured: {
      type: Boolean,
      default: false,
    },
  })

  const posts = data.filter((el) => {
    if (!props.featured) {
      if (props.tagFilter.toLowerCase() === 'none') {
        return true
      } else if (props.tagFilter.toLowerCase() === el.frontmatter.tag.toLowerCase()) {
        return true
      } else {
        return false
      }
    } else {
      if (el.frontmatter.featured) {
        return true
      } else return false
    }
  })
</script>

<template>
  <slot v-if="posts.length">
    <h2 v-if="posts.length > 0">All Posts</h2>
  </slot>
  <section class="card-container">
    <Post
      :image="post.frontmatter.image"
      :title="post.frontmatter.title"
      :tag="post.frontmatter.tag"
      :url="post.url"
      :date="post.frontmatter.date"
      :author="post.frontmatter.author"
      v-for="post in posts"
    />
  </section>
</template>

<style scoped>
  .card-container {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    flex-direction: row;
  }
</style>
