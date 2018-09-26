<template>
  <div class="layout_container">
    <section class="header">
      <nuxt-link
        to="/"
        class="header_title_text">kik4</nuxt-link>
    </section>
    <section class="main">
      <nuxt/>
    </section>
    <section class="footer">
      <div class="footer_tags">
        <div>サイトのリポジトリ：<a href="https://github.com/kik4/my-portfolio-nuxt">Github</a></div>
        <div>デプロイ：CiecleCI</div>
        <div>ホスト：Firebase</div>
      </div>
      <div>
        ©2018 kik4, All Rights Reserved.
      </div>
    </section>
  </div>
</template>

<style>
html {
  font-family: "Noto Sans JP", sans-serif;
  font-size: 15px;
}
.layout_container {
  margin: auto;
}
.main {
  background: white;
  color: #444444;
  padding: 5px 0;
}
.header {
  background: #999932;
  padding: 5px 15px;
  margin-bottom: 10px;
}
.header_title_text,
.header_title_text:hover {
  color: white;
  text-decoration: none;
  font-weight: bold;
  font-size: 32px;
}
/* Responsive */
@media (min-width: 600px) {
  .layout_container {
    width: 80%;
    max-width: 1080px;
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
  padding: 5px 15px;
  color: white;
  font-size: 0.8rem;
}
.footer a {
  color: white;
  text-decoration: underline;
}
.footer_tags {
  display: flex;
  flex-wrap: wrap;
}
.footer_tags > div {
  margin-right: 8px;
  white-space: nowrap;
}
</style>

