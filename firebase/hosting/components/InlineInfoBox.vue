<template>
  <div class="box">
    <slot />
  </div>
</template>

<style scoped>
.box {
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 6.25px;
}
.box > * {
  padding-right: 10px;
}
</style>
