<template>
  <div class="container">
    <h1 class="page_title">Virtual Scroll</h1>
    <div
      ref="vscroll"
      class="vscroll">
      <div
        :style="{transform: 'translateY(' + (start * 50) +'px)'}"
        class="vscroll_container">
        <div
          v-for="item in vlist"
          :key="item"
          class="vscroll_item">
          Item #{{ item }}
        </div>
      </div>
      <div
        :style="{transform: 'scaleY(' + (list.length * 50) +')'}"
        class="vscroll_spacer"/>
    </div>
    <p>表示される要素のみをレンダリングすることで要素が大量にあってもスクロールの挙動を軽くできます。</p>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "Virtual Scroll",
      meta: [
        {
          hid: "og:title",
          property: "og:title",
          content: "Virtual Scroll | kik4.work",
        },
      ],
    }
  },
  data() {
    return {
      list: [...Array(10000).keys()],
      scrollTop: 0,
      start: 0,
    }
  },
  computed: {
    vlist() {
      const count = 200 / 50 + 6
      return this.list.slice(this.start, this.start + count)
    },
  },
  mounted() {
    this.$refs.vscroll.addEventListener("scroll", this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.scrollTop = this.$refs.vscroll.scrollTop
      const tmpstart = Math.floor(this.scrollTop / 50)
      this.start = tmpstart - 3 > 0 ? tmpstart - 3 : 0
    },
  },
}
</script>

<style scoped>
.vscroll {
  height: 200px;
  width: 200px;
  border: black 1px solid;
  overflow: auto;
  position: relative;
}
.vscroll_container {
  position: absolute;
  top: 0;
  left: 0;
}
.vscroll_item {
  height: 50px;
}
.vscroll_spacer {
  position: absolute;
  top: 0;
  left: 0;
  height: 1px;
  width: 1px;
  transform-origin: 0 0;
}
</style>
