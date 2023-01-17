<template>
  <page-container>
    <section style="display: flex; height: 100%; flex-direction: column;">
      <el-page-header title="链接管理">
        <template #icon>
          <el-icon>
            <connection-icon/>
          </el-icon>
        </template>
        <template #content>
          <el-button type="primary" @click="openConnectionUpdatePanelHandler(connectionUpdatePanelRef)">添加</el-button>
        </template>
      </el-page-header>
      <el-divider/>
      <el-table height="100%" v-loading="loading" :show-header="false" :data="list">
        <el-table-column>
          <template #default="{row}">
            <redis-cell :data="row"/>
          </template>
        </el-table-column>
        <el-table-column width="200" align="center">
          <template #default="{row}">
            <el-button @click="infoHandler(row)" type="primary" circle :icon="InfoFilled"></el-button>
            <el-button type="success" @click="openConnectionUpdatePanelHandler(connectionUpdatePanelRef ,row.id)" circle
                       :icon="EditPen"></el-button>
            <el-button type="danger" circle :icon="Delete"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <connection-update-panel ref="connectionUpdatePanelRef" @update-success="onUpdateSuccessHandler"/>
    </section>
  </page-container>
</template>

<script lang="ts" setup>
import PageContainer from "@/components/layout/PageContainer.vue";
import ConnectionIcon from "@/icon/ConnectionIcon.vue";
import {ref} from "vue";
import ConnectionUpdatePanel from "@/components/connection/ConnectionUpdatePanel.vue";
import RedisCell from "@/components/redis/RedisCell.vue";
import {Delete, EditPen, InfoFilled} from "@element-plus/icons-vue";

import {ConnectionUpdatePanelType, useConnection} from '@/views/connection'

const connectionUpdatePanelRef = ref<ConnectionUpdatePanelType>(null)


const {
  list,
  loading,
  infoHandler,
  onUpdateSuccessHandler,
  openConnectionUpdatePanelHandler
} = useConnection()



</script>

<style lang="scss" scoped>
</style>