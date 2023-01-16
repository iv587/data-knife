<template>
  <el-container>
    <el-aside style="width: auto ;height: 100vh; display: flex; flex-direction: column; justify-content: space-between">
      <dk-menu :menu-list="list"/>
    </el-aside>
    <el-container class="dk-container">
      <el-header class="dk-header">
        <dk-header/>
      </el-header>
      <el-main style="padding: 10px; height: 100%; width: 100%; background-color: var(--el-bg-color-page)">
        <div style="height: 100%; width: 100%">
          <router-view :key="route.fullPath"/>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
import DkHeader from '@/components/layout/DkHeader.vue'
import {useMenuData} from "@/views/main/menu";

const {list , route} = useMenuData()

</script>


<style lang="scss" scoped>
.dk-container {
  width: 100%;
  height: 100vh;
}

.dk-header {
  flex-shrink: 0;
  height: $headerHeight;
  padding: 0;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--el-menu-border-color);
}
</style>