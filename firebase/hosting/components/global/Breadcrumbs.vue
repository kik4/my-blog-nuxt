<template>
  <div class="breadcrumbs">
    <template v-for="(item, index) in vlist">
      <span :key="item.path">
        <nuxt-link
          v-if="item.path"
          :to="item.path"
          class="link elem">{{ item.text }}</nuxt-link>
        <span
          v-else
          class="elem">{{ item.text }}</span>
        <span v-if="vlist.length - 1 > index"> &gt; </span>
        <span v-else/>
      </span>
    </template>
  </div>
</template>

<script>
export default {
  props: {
    list: {
      type: Array,
      required: true,
      validator(arr) {
        return arr.every(val => val.text)
      },
    },
  },
  computed: {
    vlist() {
      const arr = [{ path: "/", text: "Top" }].concat(this.list)
      const tmplist = JSON.parse(JSON.stringify(arr)) // deep copy
      tmplist[tmplist.length - 1].path = ""
      return tmplist
    },
  },
}
</script>

<style scoped>
.breadcrumbs {
  font-size: 0.8rem;
  color: #d5d5d5;
  padding-bottom: 0.5rem;
  border-bottom: #888 dotted 1px;
  margin-bottom: 0.8rem;
}
.elem {
  color: #888;
}
.link {
  text-decoration: #888 underline;
}
.link:hover {
  text-decoration: none;
}
</style>
