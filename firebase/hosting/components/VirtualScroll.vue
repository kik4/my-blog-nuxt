<template>
  <div
    ref="vscroll"
    class="vscroll">
    <div
      :style="{transform: 'translateY(' + (start * itemSize) +'px)'}"
      class="vscroll_container">
      <div
        v-for="item in vlist"
        :key="item">
        <slot :item="item"/>
      </div>
    </div>
    <div
      :style="{transform: 'scaleY(' + (items.length * itemSize) +')'}"
      class="vscroll_spacer"/>
  </div>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
    "item-size": {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      scrollTop: 0,
      start: 0,
    }
  },
  computed: {
    vlist() {
      const count = 200 / this.itemSize + 7
      return this.items.slice(this.start, this.start + count)
    },
    height() {
      return 0
    },
  },
  mounted() {
    this.$refs.vscroll.addEventListener("scroll", this.handleScroll)
  },
  unmounted() {
    this.$refs.vscroll.removeEventListener("scroll", this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.scrollTop = this.$refs.vscroll.scrollTop
      const tmpstart = Math.floor(this.scrollTop / this.itemSize)
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
.vscroll_spacer {
  position: absolute;
  top: 0;
  left: 0;
  height: 1px;
  width: 1px;
  transform-origin: 0 0;
}
</style>
