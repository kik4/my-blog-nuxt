<template>
  <div class="container">
    <Breadcrumbs :list="breadcrumbs" />
    <h1 class="page_title">Sample</h1>
    <div class="article">
      <div class="date">2018/10/31</div>
      <h2><nuxt-link to="/sample/vscroll">Virtual Scroll</nuxt-link></h2>
      <div class="body">仮想スクロール</div>
    </div>
  </div>
</template>

<script>
const blist = [{ path: "/sample", text: "Sample" }]

export default {
  head() {
    return {
      title: "Sample",
      meta: [
        {
          hid: "og:title",
          property: "og:title",
          content: "Sample | kik4.work",
        },
      ],
      __dangerouslyDisableSanitizers: ["script"],
      script: [this.jsonldBreadcrumbs(blist)],
    }
  },
  computed: {
    breadcrumbs() {
      return blist
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
