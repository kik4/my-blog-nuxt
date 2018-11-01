<template>
  <div class="container">
    <h1 class="page_title">Sample</h1>
    <div class="article">
      <div class="date">2018/10/31</div>
      <h2><nuxt-link to="/sample/vscroll">Virtual Scroll</nuxt-link></h2>
      <div class="body">仮想スクロール</div>
    </div>
  </div>
</template>

<script>
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
      script: [
        {
          hid: "jsonld",
          innerHTML: `{
            "@context": "http://schema.org",
            "@type": "BreadcrumbList",
            "itemListElement": [
              {
                "@type": "ListItem",
                "position": 1,
                "item": {
                  "@id": "${this.base_url()}",
                  "name": "kik4.work"
                }
              }, {
                "@type": "ListItem",
                "position": 2,
                "item": {
                  "@id": "${this.url()}",
                  "name": "Sample"
                }
              }
            ]
          }`,
          type: "application/ld+json",
        },
      ],
    }
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
