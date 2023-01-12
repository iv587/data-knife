<template>
<section class="web-container">
  <slot></slot>
</section>
</template>

<script lang="ts" setup>
</script>

<style lang="scss">
@media screen and (max-width: 600px){
  .web-container {
    position: absolute;
    width: 100%;
    height: 100%;
  }
}

@media screen and (min-width: 800px) {
  .web-container {
    width: 100vw;
    height: 100vh;
  }
}
</style>