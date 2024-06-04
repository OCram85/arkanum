<script setup>
  import Posts from './Posts.vue'
</script>

<template>
  <div class="blog-container">
    <h1>Blog</h1>
    <h2>⭐ Featured Articles</h2>
    <Posts featured/>
    <h2>📰 All Posts</h2>
    <Posts/>
  </div>
</template>

<style scoped>
  .blog-container {
    margin: auto;
    width: 100%;
    max-width: 1280px;
    padding: 0 64px;

  }

  h1, h2, h3 {
    position: relative;
    font-weight: 600;
    outline: none;
  }

  h1 {
    letter-spacing: -0.02em;
    line-height: 40px;
    font-size: 32px;
    margin-top: 32px;
  }

  h2 {
    margin: 40px 0 16px;
    border-top: 1px solid var(--vp-c-divider);
    padding-top: 24px;
    letter-spacing: -0.02em;
    line-height: 32px;
    font-size: 24px;
  }

  h3 {
    margin: 32px 0 0;
    letter-spacing: -0.01em;
    font-size: 20px;
  }
</style>
