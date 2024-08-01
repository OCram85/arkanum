<script setup>
import { withBase } from 'vitepress';
import { computed } from 'vue';

const props = defineProps({
  Source: {
    type: String,
    required: true
  },
  Caption: {
    type: String
  }
})

const ImageSource = computed(() => {
  return withBase(props.Source)
})
</script>

<template>
  <div class="blog-header-container">
    <img :src="ImageSource">
    <p v-if="props.Caption">{{ props.Caption }}</p>
  </div>
</template>

<style scoped>
.blog-header-container {
  width: 100%;
  margin: 24px 0;
}

.blog-header-container img {
  max-width: 325px;
  max-height: 325px;
  margin: 12px auto;
  border: 2px solid var(--vp-c-brand-soft);
  border-radius: 12px;
  /*box-shadow: -1px 5px 24px 12px var(--vp-c-brand-soft);*/
  box-shadow: 0px 8px 16px var(--vp-c-brand-soft);
}

.blog-header-container p {
  width: 80%;
  margin: 0 auto;
  text-align: center;
  font-size: smaller;
}
</style>
