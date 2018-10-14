<template>
  <div class="menu">
    <template v-if="$nuxt.$route.name !== path">
      <nuxt-link
        :to="path"
        class="menu_a">{{ path }}</nuxt-link>
    </template>
    <template v-else>
      <span class="menu_s">{{ path }}</span>
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
  },
}
</script>

<style scoped>
.menu > * {
  margin-left: 5px;
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
