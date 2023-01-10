<template>
  <el-container>
    <dk-side/>
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
import DkSide from '@/components/layout/DkSide.vue'
import DkHeader from '@/components/layout/DkHeader.vue'
import {useRoute, useRouter} from "vue-router"
import {watch} from "vue";
import useMenuStore from "@/store/menu"

const route = useRoute()
const router = useRouter()
const currentRoute =  router.currentRoute
const menuStore = useMenuStore()

watch(currentRoute, (newRoute, oldRoute) => {
  menuStore.setActiveIndex(newRoute.path)
})

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