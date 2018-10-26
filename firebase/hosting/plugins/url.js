import Vue from "vue"

Vue.mixin({
  methods: {
    base_url() {
      return process.env.BASE_URL
    },
    url() {
      let url = process.env.BASE_URL + this.$route.fullPath
      if (url.slice(-1) === "/") {
        url = url.slice(0, -1)
      }
      return url
    },
  },
})
