<template>
  <el-icon class="icon-button" :size="size">
    <slot></slot>
  </el-icon>
</template>

<script lang="ts" setup>
defineProps<{size?: number; color?: string}>()
</script>

<style lang="scss">
.icon-button {
  color: var(--el-color-primary);
  &:hover {
    cursor: pointer;
    color: var(--el-color-primary-dark-2);

  }
}
</style>