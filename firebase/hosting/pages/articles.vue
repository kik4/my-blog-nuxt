<template>
  <div class="container">
    <Breadcrumbs :list="breadcrumbs" />
    <h1 class="page_title">Articles</h1>
    <template v-if="items && items.length > 0">
      <div
        v-for="item in items"
        :key="item.id"
        class="article"
      >
        <div class="date">{{ jdate(item.created_at) }} <a href="https://qiita.com/kik4">Qiita</a></div>
        <h2><a :href="item.url">{{ item.title }}</a></h2>
        <div class="body">{{ item.body.substring(0, 100) }}...</div>
      </div>
    </template>
    <template v-else>
      <div
        v-for="item in [1,2,3,4,5,6]"
        :key="item"
        class="article skelton"
      >
        <span class="date">&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;</span>
        <h2>
          <span>&nbsp;</span>
        </h2>
        <div class="body">&nbsp;</div>
      </div>
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
    if (this.$store.state.blog.articles && this.$store.state.blog.articles.length > 0) {
      return
    }
    this.$store.dispatch("blog/getArticles")
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
  margin-bottom: 2px;
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
.skelton > * {
  color: transparent;
  background: linear-gradient(60deg, #e0e0e0, #e0e0e0, #eeeeee, #e0e0e0, #e0e0e0);
  background-size: 500%;
  animation: SkeltonAnim 2s linear infinite;
}
@keyframes SkeltonAnim {
  0% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}
</style>
