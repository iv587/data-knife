<template>
<div class="menu-logo">
  <p class="menu-logo-text" v-if="collapse">D K</p>
  <p style="flex-shrink: 0" class="menu-logo-text" v-else >Data Knife</p>
</div>
</template>

<script lang="ts" setup>
import MenuStore from '@/store/menu'
import {computed} from "vue";

const menuStore = MenuStore()

const collapse = computed(() => {
  return menuStore.collapse
})

</script>

<style lang="scss" scoped>
.menu-logo {
  
  flex-shrink: 0;
  height: $headerHeight;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20pt;
  color: var(--el-color-primary-dark-2);
  font-family: "阿里巴巴普惠体 2.0 115 Black";
}
</style>