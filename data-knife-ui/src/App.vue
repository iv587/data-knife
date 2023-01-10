<script setup lang="ts">
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import MainView from "@/views/MainView.vue";

</script>

<template>
  <el-config-provider :locale="zhCn">
    <router-view />
  </el-config-provider>
</template>

<style scoped>
</style>
