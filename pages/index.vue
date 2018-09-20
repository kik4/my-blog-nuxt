<template>
  <div>
    <h1 class="title">トップページ</h1>
    <p><nuxt-link to="counter">カウンタへ</nuxt-link></p>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      count: 0
    }
  },
  head: {
    titleTemplate: "kik4"
  }
}
</script>

<style>
</style>

