<template>
  <div class="menu">
    <a
      :href="path"
      class="menu_a">{{ text }}</a>
  </div>
</template>

<script>
export default {
  props: {
    path: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  },
}
</script>

<style scoped>
.menu {
  font-size: 1.2rem;
  text-decoration: underline;
}
.menu > * {
  margin-right: 5px;
  font-weight: 600;
}
.menu_a {
  color: rgb(255, 200, 200);
}
.menu_a:hover {
  color: rgb(255, 95, 95);
}
</style>
