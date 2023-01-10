<template>
  <el-aside style="width: auto ;height: 100vh; display: flex; flex-direction: column; justify-content: space-between">
    <dk-menu :menu-list="list"/>
  </el-aside>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import useMenuStore from "@/store/menu";

const menuStore = useMenuStore()

const list = computed(() => {
  return menuStore.menuList
})

</script>

<style scoped>

</style>