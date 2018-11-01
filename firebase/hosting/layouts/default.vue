<template>
  <div class="layout_container">
    <section class="header">
      <div class="header_image">
        <nuxt-link
          to="/"
          class="header_title_text">kik4.work</nuxt-link>
        <div class="menu">
          <InternalHeaderMenu
            path="/articles"
            text="Articles"/>
          <InternalHeaderMenu
            path="/sample"
            text="Sample"/>
          <ExternalHeaderMenu
            path="https://kik4.hatenablog.com"
            text="Blog" />
          <ExternalHeaderMenu
            path="https://twitter.com/_kik4_"
            text="Twitter" />
          <ExternalHeaderMenu
            path="https://github.com/kik4"
            text="GitHub" />
          <ExternalHeaderMenu
            path="https://qiita.com/kik4"
            text="Qiita" />
        </div>
      </div>
    </section>
    <section class="main">
      <nuxt/>
    </section>
    <section class="footer">
      <div class="footer_menu">
        <span>Sitemap :</span>
        <nuxt-link to="/">Top</nuxt-link>
        <nuxt-link to="/articles">Articles</nuxt-link>
        <nuxt-link to="/sample">Sample</nuxt-link>
      </div>
      <div class="footer_menu">
        <span>Link :</span>
        <a href="https://kik4.hatenablog.com">Blog</a>
        <a href="https://twitter.com/_kik4_">Twitter</a>
        <a href="https://github.com/kik4">GitHub</a>
        <a href="https://qiita.com/kik4">Qiita</a>
      </div>
      <div>
        ©2018 kik4, All Rights Reserved.
      </div>
    </section>
  </div>
</template>

<script>
export default {
  head() {
    return {
      meta: [{ hid: "og:url", property: "og:url", content: this.url(this.$route) }],
    }
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
  clip-path: polygon(0% 0%, 100% 0%, 100% 82%, 15% 90%, 9% 100%, 10% 90%, 0% 82%);
  padding: 5px 0 25px 15px;
}
.header_title_text,
.header_title_text:hover {
  color: white;
  text-decoration: none;
  font-weight: bold;
  font-size: 2rem;
}
.menu {
  display: flex;
  flex-wrap: wrap;
  align-items: flex-end;
}
/* Responsive */
@media (min-width: 600px) {
  html {
    font-size: 15px;
  }
  .layout_container {
    width: 80%;
    max-width: 1080px;
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
  color: lightgray;
  font-size: 0.9rem;
}
.footer_menu {
  margin-bottom: 3px;
}
.footer_menu > span {
  color: wheat;
  margin-right: 7px;
}
.footer_menu > a {
  color: white;
  margin-right: 5px;
}
</style>
