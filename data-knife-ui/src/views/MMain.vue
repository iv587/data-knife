<template>
  <el-container>
    <el-header class="dk-header">
      <dk-header/>
    </el-header>
    <el-main style="padding: 10px; height: 100%; width: 100%; background-color: var(--el-bg-color-page)">
      <div style="height: 100%; width: 100%">
        <router-view :key="route.fullPath"/>
      </div>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'MobileMainView',
}
</script>

<style scoped>

</style>