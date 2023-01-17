<template>
  <section class="fill" style="display: flex; flex-direction: column;">
    <section class="header">
      <el-page-header>
        <template #title><span></span></template>
        <template #icon>
          <el-select @change="onSelect" v-model="selectDbNum">
            <el-option :value="item" :label="`DB${item}`" v-for="item in dbList">
              <template #default>
                <el-icon>
                  <redis-db-icon/>
                </el-icon>
                <span>DB{{ item }}</span>
              </template>
            </el-option>
          </el-select>
        </template>
        <template #content>
          <el-space>
            <el-button type="primary" plain @click="openUpdatePanelHandler(redisDataUpdatePanel)">添加</el-button>
            <el-button type="success" plain @click="showDashboardHandler(emit,true)">面板</el-button>
          </el-space>
        </template>
        <el-input @clear="listKeysHandler" clearable v-model="keyPattern" placeholder="输入key进行搜索">
          <template #prefix>
            <el-checkbox :true-label="1" :false-label="0" v-model="pattern.left">*</el-checkbox>
          </template>
          <template #suffix>
            <el-checkbox :true-label="1" :false-label="0" v-model="pattern.right">*</el-checkbox>
          </template>
          <template #append>
            <el-button @click="listKeysHandler" :icon="Search"/>
          </template>
        </el-input>
        <el-alert :closable="false">共有key{{ keysRes.total }}个，本次获取{{ keysRes.max }}</el-alert>
      </el-page-header>
    </section>
    <el-table :show-header="false" v-loading="dataFetchLoading" :data="keysRes.list" border max-height="100%">
      <el-table-column>
        <template #default="{row}">
          <section class="redis-m-cell">
            <section class="cell-left">
              <h3>{{ row.key }}</h3>
              <section class="footer">
                <el-space>
                  <el-tag effect="plain" :type="dataTypeMap.get(row.type)">{{ row.type }}</el-tag>
                  <el-tag v-if="row.ttl > 0" effect="plain" type="info">{{ row.ttl }}后秒过期</el-tag>
                </el-space>
              </section>
            </section>
            <section class="cell-right">
              <el-space>
                <el-button @click="openUpdatePanelHandler(redisDataUpdatePanel, row.key)" circle :icon="EditPen" type="success"></el-button>
                <el-button circle :icon="Delete" type="danger"></el-button>
              </el-space>
            </section>
          </section>
        </template>
      </el-table-column>
    </el-table>
    <redis-data-update-panel width="95%" :connection-id="id" :db-no="selectDbNum" @update-success="listKeysHandler"
                             ref="redisDataUpdatePanel"/>
  </section>
</template>

<script lang="ts" setup>
import RedisDb from "@/components/redis/RedisDb.vue";
import {Delete, EditPen, Search} from "@element-plus/icons-vue";
import RedisDbIcon from "@/icon/RedisDbIcon.vue";
import {ref} from "vue";
import RedisDataUpdatePanel from "@/components/redis/RedisDataUpdatePanel.vue";
import {RedisDataUpdatePanelType, useRedisData} from '@/views/redis/redisData'
const redisDataUpdatePanel = ref<RedisDataUpdatePanelType>(null)
const props = defineProps<{
  id: number,
}>()
const {
  pattern, keyPattern, dbList, selectDbNum, dataFetchLoading,
  keysRes, dataTagTypeMap, onSelect, listKeysHandler
  , showDashboardHandler, openUpdatePanelHandler
} = useRedisData(props.id)

const dataTypeMap = ref<Map<string, string>>(dataTagTypeMap)

const emit = defineEmits<{
  (e: 'close'): void
}>()
</script>

<style lang="scss" scoped>
.redis-m-cell {
  display: flex;
  align-items: center;
  justify-content: space-between;
  .cell-left {
    display: flex;
    flex-direction: column;
    h3 {
      color: var(--el-color-info);
    }
    .footer {
      display: flex;
      align-items: center;
    }
  }
  .cell-right {

  }



}


</style>