<template>
  <div class="container">
    <section class="header">
      <p>
        <a
          href="/"
          class="header_title_text">kik4</a>
      </p>
    </section>
    <section class="main">
      <nuxt/>
    </section>
  </div>
</template>

<style>
html {
  font-family: "Noto Sans JP", sans-serif;
}
.container {
  margin: auto;
}
.main {
  background: white;
  color: #222222;
}
.header {
  background: #222222;
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
</style>

