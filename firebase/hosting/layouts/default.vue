<template>
  <div class="layout_container">
    <section class="header">
      <div class="header_image">
        <nuxt-link
          to="/"
          class="header_title_text">kik4</nuxt-link>
        <InternalHeaderMenu
          path="articles"
          text="Articles"/>
        <ExternalHeaderMenu
          path="https://kik4.hatenablog.com"
          text="Blog" />
        <ExternalHeaderMenu
          path="https://twitter.com/_kik4_"
          text="Twitter" />
        <ExternalHeaderMenu
          path="https://github.com/kik4"
          text="GitHub" />
      </div>
    </section>
    <section class="main">
      <nuxt/>
    </section>
    <section class="footer">
      ©2018 kik4, All Rights Reserved.
    </section>
  </div>
</template>

<script>
export default {
  head() {
    return {
      meta: [{ hid: "og:url", property: "og:url", content: this.url() }],
    }
  },
  methods: {
    url() {
      let url = process.env.baseUrl + this.$route.fullPath
      if (url.slice(-1) === "/") {
        url = url.slice(0, -1)
      }
      return url
    },
  },
}
</script>

<style>
html {
  font-family: "Noto Sans JP", sans-serif;
  font-size: 12.5px;
}
body {
  background: #eeeeee;
  overflow-y: scroll; /* ガタつき防止 */
}
.layout_container {
  margin: auto;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.2);
}
.main {
  background: white;
  color: #444444;
  padding: 10px 4px 4px 4px;
}

.header {
  background: white;
}
.header_image {
  background: #999932;
  clip-path: polygon(0% 0%, 100% 0%, 100% 75%, 15% 80%, 9% 100%, 10% 80%, 0% 75%);
  padding: 5px 0 25px 15px;
  display: flex;
  flex-wrap: wrap;
  align-items: flex-end;
}
.header_title_text,
.header_title_text:hover {
  color: white;
  text-decoration: none;
  font-weight: bold;
  font-size: 32px;
  line-height: 32px;
  margin-right: 10px;
}
/* Responsive */
@media (min-width: 600px) {
  .layout_container {
    width: 80%;
    max-width: 1080px;
    font-size: 15px;
  }
  .main {
    padding: 12px;
  }
}

.page_title {
  border-color: #999932;
  border-width: 0 0 1px 3px;
  border-style: solid;
  padding: 0 0 3px 7px;
}

.footer {
  background: #999932;
  padding: 15px;
  color: white;
  font-size: 0.8rem;
}
</style>
