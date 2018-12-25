<template>
  <div
    ref="vscroll"
    class="vscroll"
  >
    <div
      :style="{transform: 'translateY(' + (start * itemSize) +'px)'}"
      class="vscroll_container"
    >
      <slot
        v-for="item in vlist"
        :item="item"
      />
    </div>
    <div
      :style="{transform: 'scaleY(' + (items.length * itemSize) +')'}"
      class="vscroll_spacer"
    />
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
      height: 0,
    }
  },
  computed: {
    vlist() {
      const count = this.height / this.itemSize + 7
      return this.items.slice(this.start, this.start + count)
    },
  },
  mounted() {
    this.$refs.vscroll.addEventListener("scroll", this.handleScroll)
    this.height = this.$refs.vscroll.clientHeight
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
  overflow: auto;
  position: relative;
  will-change: scroll-position;
  transform: translateZ(0);
  -webkit-overflow-scrolling: touch;
}
.vscroll_container {
  position: absolute;
  top: 0;
  left: 0;
  min-width: 100%;
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
