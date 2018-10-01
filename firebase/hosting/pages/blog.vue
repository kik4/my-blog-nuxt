<template>
  <div class="container">
    <h1 class="page_title">Blog</h1>
    <template v-if="items!==null">
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
  head() {
    return {
      title: "blog",
      meta: [{ hid: "og:title", property: "og:title", content: "blog | kik4" }],
    }
  },
  data: function() {
    return {
      items: [],
    }
  },
  created() {
    axios
      .get("https://qiita.com/api/v2/users/kik4/items")
      .then(
        function(response) {
          this.items = response.data
        }.bind(this),
      )
      .catch(
        function() {
          this.items = null
        }.bind(this),
      )
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

