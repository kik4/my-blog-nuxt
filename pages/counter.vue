<template>
  <div>
    <h1 class="title">カウンタ</h1>
    <p>
      特に意味のないカウンタ: {{ count }}<br>
      <input
        type="button"
        value="Click!"
        @click="count++">
    </p>
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
    title: "カウンタ"
  }
}
</script>

<style>
</style>

