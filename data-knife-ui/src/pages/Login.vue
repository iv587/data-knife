<template>
  <el-config-provider :locale="zhCn">
    <login-view/>
  </el-config-provider>
</template>

<script lang="ts" setup>
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import LoginView from '@/pages/LoginView.vue'</script>

<style scoped>

</style>