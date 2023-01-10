<template>
<section class="page-container">
  <section class="container">
    <slot></slot>
  </section>

</section>
</template>

<script lang="ts" setup>
</script>

<style scoped lang="scss">
.page-container {
  background-color: var(--el-bg-color);
  width: calc(100% - 20px);
  height: calc(100% - 20px);
  #boxShadow: var(--el-box-shadow-light);
  #border: 1px solid var(--el-border-color);
  #border-radius: var(--el-border-radius-base);
  padding: 10px;
  .container {
    width: 100%;
    height: 100%;
  }
}
</style>