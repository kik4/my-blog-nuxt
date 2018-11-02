<template>
  <div class="container">
    <Breadcrumbs :list="breadcrumbs" />
    <h1 class="page_title">Virtual Scroll</h1>
    <VirtualScroll
      :items="items"
      :item-size="50"
      class="vscroll">
      <div
        slot-scope="{item}"
        :style="{'background-color': `hsl(${item * 20 % 360}, 100%, 90%)`}"
        :key="item"
        class="item"
      >Item #{{ item }}</div>
    </VirtualScroll>
    <p>表示される要素のみをレンダリングすることで要素が大量にあってもスクロールの挙動を軽くできます。</p>
  </div>
</template>

<script>
import VirtualScroll from "@/components/VirtualScroll"

const blist = [{ path: "/sample", text: "Sample" }, { path: "/sample/vscroll", text: "Virtual Scroll" }]

export default {
  components: { VirtualScroll },
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
      __dangerouslyDisableSanitizers: ["script"],
      script: [this.jsonldBreadcrumbs(blist)],
    }
  },
  data() {
    return {
      items: [...Array(10000).keys()],
    }
  },
  computed: {
    breadcrumbs() {
      return blist
    },
  },
}
</script>

<style scoped>
.vscroll {
  height: 400px;
  width: 400px;
  border: black 1px solid;
}
.item {
  height: 50px;
}
</style>
