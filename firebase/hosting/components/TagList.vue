<template>
  <ul
    v-if="list"
    class="tags">
    <li
      v-for="item in list"
      :key="item">{{ item }}</li>
  </ul>
</template>

<script>
export default {
  props: {
    list: {
      type: Array,
      default: () => [],
    },
  },
}
</script>

<style scoped>
.tags {
  display: flex;
  flex-wrap: wrap;
  padding-left: 0;
}
.tags > li {
  list-style: none;
  margin-right: 2px;
  white-space: nowrap;
  border: 1px solid #cccccc;
  border-radius: 5px;
  padding: 0 5px;
  background-color: #fff9f9f9;
}
</style>
