<template>
  <el-sub-menu  v-if="menuItem.children && menuItem.children.length > 0" :index="menuItem.index">
    <template #title>
      <el-icon v-if="menuItem.icon">
        <component :is="menuItem.icon" />
      </el-icon>
      <span>{{menuItem.title}}</span></template>
    <dk-menu-item v-for="item in menuItem.children" :menu-item="item" :key="item.index" />
  </el-sub-menu>
  <el-menu-item v-else :index="menuItem.index">
    <el-icon v-if="menuItem.icon">
      <component :is="menuItem.icon" />
    </el-icon>
    <span>{{menuItem.title}}</span>
  </el-menu-item>
</template>

<script lang="ts" setup>

import {MenuItem} from "@/type";
import {useRouter} from 'vue-router'
import useMenuStore from "@/store/menu";


const menuStore = useMenuStore()


const props = defineProps<{menuItem: MenuItem}>()


</script>

<style scoped>

</style>