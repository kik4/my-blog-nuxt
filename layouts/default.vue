<template>
  <div class="container">
    <section class="header">
      <nuxt-link
        to="/"
        class="header_title_text">kik4</nuxt-link>
    </section>
    <section class="main">
      <nuxt/>
    </section>
  </div>
</template>

<style>
html {
  font-family: "Noto Sans JP", sans-serif;
  font-size: 15px;
}
.container {
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
  .container {
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
</style>

