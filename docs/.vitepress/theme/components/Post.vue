<script setup>
  import { computed } from 'vue'
  import { withBase } from 'vitepress'
  import Badge from 'vitepress/dist/client/theme-default/components/VPBadge.vue'

  const props = defineProps({
    image: {
      type: String,
      default: 'placeholder.jpg',
    },
    title: {
      type: String,
      required: true,
    },
    tag: {
      type: String,
      default: 'no-tag',
    },
    url: {
      type: String,
      required: true,
    },
    featured: {
      type: Boolean,
      default: false,
    },
    date: {
      type: String,
      default: '2020-01-01'
    },
    author: {
      type: String,
      default: 'Anonymous'
    }
  })

  const normTitle = computed(() => {
    if (props.title.length > 50) {
      return props.title.slice(0, 50) + '...'
    } else {
      return props.title
    }
  })

  const normImage = computed(() => {
    if (props.image !== null) return withBase(props.image)
    else {
      return withBase('placeholder.jpg')
    }
  })

  const normTag = computed(() => {
    if (props.tag !== null) return props.tag.toUpperCase()
    else {
      return 'EMPTY-TAG'
    }
  })

const authorURL = computed(() => {
    return `https://gitea.ocram85.com/${props.author}`
  })
</script>

<template>
  <article class="card">
    <div class="card-header">
      <img :src="normImage" />
    </div>
    <div class="card-content">
      <div class="badge-container">
        <Badge :text="normTag" type="tip" />
      </div>
      <h3>{{ normTitle }}</h3>
      <p class="acredits">by <a class="author" :href="authorURL" target="_blank">{{ author }}</a> on <div class="date">{{ date}}</div></p>
      <a class="url" :href="url">Read More</a>
    </div>
  </article>
</template>

<style scoped>
  .card {
    margin: 15px;
    padding: 0px;
    background-color: var(--vp-c-bg-soft);
    border-color: var(--vp-c-brand-2);
    border-style: solid;
    border-width: 1px;
    border-radius: 10px;
    width: 325px;
    /*height: 450px; */
    overflow: hidden;
    box-shadow: 0px 8px 16px var(--vp-c-brand-soft);
    position: relative;
    top: 0;
  }

  .card:hover {
    border-color: var(--vp-c-brand-1);
    filter: brightness(125%);
    transition: 0.25s;
    box-shadow: 0px 16px 32px var(--vp-c-brand-soft);
    /*top: -10px;*/
  }

  .card-header {
    width: 100%;
    height: 200px;
  }
  .card-header img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .card-content {
    padding: 0 20px;
    height: 200px;
  }

  .card-content h3 {
    margin-top: 5px;
    height: 60px;
  }

  .badge-container {
    margin: 10px 0;
  }

  h3 {
    position: relative;
    font-weight: 600;
    outline: none;

    margin: 32px 0 0;
    letter-spacing: -0.01em;
    line-height: 28px;
    font-size: 20px;
  }

  a {
    font-weight: 500;
    color: var(--vp-c-brand-1);
    text-decoration: underline;
    text-underline-offset: 2px;
    transition: color 0.025s, opacity 0.25s;
  }

  a:hover {
    color: var(--vp-c-brand-2);
  }

  .acredits {
    margin-top: 10px;
    display: inline-block;
    height: 50px;

    font-weight: 300;
    font-size: 14px;
  }

  .author {
    display: inline-block;
    font-weight: 400;
    text-decoration: none;
    color: var(--vp-c-brand-2)
  }

  .date {
    display: inline-block;
    font-weight: 400;
    color: var(--vp-c-brand-2);
  }

  .url {
    display: block;
    margin-top: 5px;
    padding-right: 30px;
    max-height: 30px;
    text-align: right;
  }
</style>
