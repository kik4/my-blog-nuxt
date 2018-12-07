<template>
  <div class="container">
    <Breadcrumbs :list="breadcrumbs" />
    <h1 class="page_title">Articles</h1>
    <template v-if="items && items.length > 0">
      <div
        v-for="item in items"
        :key="item.id"
        class="article">
        <div class="date">{{ jdate(item.created_at) }} <a href="https://qiita.com/kik4">Qiita</a></div>
        <h2><a :href="item.url">{{ item.title }}</a></h2>
        <div class="body">{{ item.body.substring(0, 100) }}...</div>
      </div>
    </template>
    <template v-else-if="items && items.length === 0">
      <div>取得中...</div>
    </template>
    <template v-else>
      <div>取得失敗</div>
    </template>
  </div>
</template>

<script>
import axios from "axios"

const blist = [{ path: "/articles", text: "Articles" }]

export default {
  async fetch({ store, params }) {
    if (!process.browser) {
      return
    }
    // check cache
    if (store.state.blog.articles && store.state.blog.articles.length > 0) {
      return
    }
    await store.dispatch("blog/getArticles")
  },
  computed: {
    items() {
      return this.$store.state.blog.articles
    },
    breadcrumbs() {
      return blist
    },
  },
  mounted() {
    // check cache
    if (store.state.blog.articles && store.state.blog.articles.length > 0) {
      return
    }
    await store.dispatch("blog/getArticles")
  },
  head() {
    return {
      title: "Articles",
      meta: [
        {
          hid: "og:title",
          property: "og:title",
          content: "Articles | kik4.work",
        },
      ],
      __dangerouslyDisableSanitizers: ["script"],
      script: [this.jsonldBreadcrumbs(blist)],
    }
  },
  methods: {
    jdate(date) {
      var date = new Date(date)
      return (
        date.getFullYear() + "/" + this.zeroPadding(date.getMonth() + 1, 2) + "/" + this.zeroPadding(date.getDate(), 2)
      )
    },
    zeroPadding(num, length) {
      return ("000000" + num).slice(-length)
    },
  },
}
</script>

<style scoped>
h2 {
  font-size: 1.2rem;
  margin-bottom: 0;
}
.article {
  padding: 5px;
}
.date {
  font-size: 0.85rem;
  font-weight: bolder;
}
.body {
  font-size: 0.8rem;
  word-wrap: break-word;
  color: rgb(153, 153, 153);
}
</style>
