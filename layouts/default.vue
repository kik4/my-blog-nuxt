<template>
  <div class="container">
    <section class="header">
      <h1>kik4's page</h1>
    </section>
    <div class="flexbox">
      <section class="side">
        <p>The Little Prince (French: Le Petit Prince), ...afafa</p>
      </section>
      <section class="main">
        <nuxt/>
      </section>
    </div>
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
  width: 70%;
}
.side {
  background: #222222;
  color: white;
  width: 30%;
}
.main,
.side {
  padding: 2%;
  border-radius: 10px;
}

/* Responsive */
@media (min-width: 600px) {
  .container {
    width: 80%;
  }
  .flexbox {
    display: -webkit-flex;
    display: flex;
  }
  .main {
    margin: 0;
  }
}
</style>

