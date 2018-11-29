<template>
  <div class="menu">
    <template v-if="$route.name !== path.replace(/^\//g, '')">
      <nuxt-link :to="path" class="menu_a">{{ text }}</nuxt-link>
    </template>
    <template v-else>
      <span class="menu_s">{{ text }}</span>
    </template>
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
}
.menu > * {
  margin-right: 5px;
  font-weight: 600;
}
.menu_a {
  color: white;
}
.menu_a:hover {
  color: lightgray;
}
.menu_s {
  color: #5b5b20;
}
</style>
