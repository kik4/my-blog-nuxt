<template>
  <div class="indent">
    <slot/>
  </div>
</template>

<style scoped>
.indent {
  padding-left: 12px;
}
@media (min-width: 600px) {
  ul {
    padding-left: 20px;
  }
}
</style>
