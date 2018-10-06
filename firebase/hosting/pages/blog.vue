<template>
  <div class="container">
    <h1 class="page_title">Blog</h1>
    <template v-if="items">
      <div
        v-for="item in items"
        :key="item.id"
        class="article">
        <div class="date">{{ jdate(item.created_at) }} Qiita</div>
        <h2><a :href="item.url">{{ item.title }}</a></h2>
      </div>
    </template>
    <template v-else>
      <div>取得失敗</div>
    </template>
  </div>
</template>

<script>
import axios from "axios"

export default {
  async fetch({ store, params }) {
    // check cache
    if (store.state.blog.articles.length) {
      return
    }
    await store.dispatch("blog/getArticles")
  },
  head() {
    return {
      title: "blog",
      meta: [{ hid: "og:title", property: "og:title", content: "blog | kik4" }],
    }
  },
  computed: {
    items() {
      return this.$store.state.blog.articles
    },
  },
  methods: {
    jdate(date) {
      var date = new Date(date)
      return date.getFullYear() + "/" + (date.getMonth() + 1) + "/" + date.getDate()
    },
  },
}
</script>

<style scoped>
h2 {
  font-size: 1.3rem;
  font-weight: normal;
  margin-bottom: 0;
}
.article {
  padding: 5px;
}
.date {
  font-size: 1rem;
}
</style>

