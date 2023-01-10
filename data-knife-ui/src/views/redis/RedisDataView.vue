<template>
  <section style="display: flex;width: 100%; height: 100%;">
    <section class="redis-view-left">
      <el-link :underline="false" type="success">数据库列表</el-link>
      <el-divider style="margin: 10px 0"/>
      <el-scrollbar style="
      border: 1px solid var(--el-border-color);
      border-radius: var(--el-border-radius-base);
      ">
        <div style="display: flex; flex-direction: column; align-items: center">
          <redis-db @select-db="onSelect" :selected="selectDbNum == item" :db-num="item" v-for="item in dbList"
                    :key="item"/>
        </div>
      </el-scrollbar>
    </section>
    <el-divider direction="vertical" style="height: 100%"/>
    <div class="redis-view-right" style="width: 100%">
      <section class="header">
        <el-page-header :title="`DB${selectDbNum}`">
          <template #icon>
            <el-icon>
              <redis-db-icon/>
            </el-icon>
          </template>
          <template #content>
            <el-space>
              <el-input @clear="listKeysHandler" clearable v-model="keyPattern" placeholder="输入key进行搜索">
                <template #prefix>
                  <el-checkbox :true-label="1" :false-label="0" v-model="pattern.left" >*</el-checkbox>
                </template>
                <template #suffix>
                  <el-checkbox :true-label="1" :false-label="0" v-model="pattern.right" >*</el-checkbox>
                </template>
                <template #append>
                  <el-button @click="listKeysHandler" :icon="Search"/>
                </template>
              </el-input>
              <el-divider direction="vertical"/>
              <el-button type="primary" plain @click="openUpdatePanelHandler()">添加</el-button>
              <el-button type="success" plain @click="showDashboardHandler(true)">面板</el-button>
            </el-space>
          </template>
          <el-alert :closable="false">共有key{{ keysRes.total }}个，本次获取{{ keysRes.max }}</el-alert>
        </el-page-header>
      </section>
      <el-table v-loading="dataFetchLoading" :data="keysRes.list" border height="100%">
        <el-table-column prop="key" label="KEY"></el-table-column>
        <el-table-column align="center" label="类型" width="150">
          <template #default="{row}">
            <el-tag effect="plain" :type="dataTypeMap.get(row.type)">{{ row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" width="150" prop="ttl" label="TTL(秒)"></el-table-column>
        <el-table-column align="center" label="操作" width="200">
          <template #default="{row}">
            <el-button @click="openUpdatePanelHandler(row.key)" circle :icon="EditPen" type="success"></el-button>
            <el-button circle :icon="Delete" type="danger"></el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </section>
  <redis-data-update-panel :connection-id="id" :db-no="selectDbNum" @update-success="listKeysHandler" ref="redisDataUpdatePanel" />
</template>

<script lang="ts" setup>
import RedisDb from "@/components/redis/RedisDb.vue";
import {Delete, EditPen, Search} from "@element-plus/icons-vue";
import RedisDbIcon from "@/icon/RedisDbIcon.vue";
import {Component, onMounted, ref} from "vue";
import {listConnection} from "@/api/connection";
import {redisDbsApi, redisKeysApi} from "@/api/redis";
import {RedisKeyListRes} from "@/type";
import RedisDataUpdatePanel from "@/components/redis/RedisDataUpdatePanel.vue";

const props = defineProps<{
  id: number,
}>()

const redisDataUpdatePanel = ref()

const pattern = ref<{
  left: number
  right: number
}>({
  left: 1,
  right: 1
})

const keyPattern = ref('')


const dbList = ref<number[]>()

const selectDbNum = ref<number>(0)

const dataFetchLoading = ref(false)

const keysRes = ref<RedisKeyListRes>({
  list: [],
  total: 0,
  max: 0
})

const dataTagTypeMap = new Map<string, string>()
dataTagTypeMap.set('list', 'waring')
dataTagTypeMap.set('set', 'info')
dataTagTypeMap.set('hash', '')
dataTagTypeMap.set('zset', 'danger')
dataTagTypeMap.set('string', 'success')

const dataTypeMap = ref<Map<string,string>>(dataTagTypeMap)


const redisDbHandler = async () => {
  const res = await listConnection()
  if (res.data.list.length > 0) {
    const dbsRes = await redisDbsApi({id: props.id})
    dbList.value = dbsRes.data.list
    selectDbNum.value = dbList.value[0]
    await listKeysHandler()
  }
}

const onSelect = async (dbNum: number) => {
  selectDbNum.value = dbNum
  await listKeysHandler()
}

const listKeysHandler = async () => {
  dataFetchLoading.value = true
  try {
    let searchKey = keyPattern.value
    const {left, right} = pattern.value
    if (left == 1) {
      searchKey = '*' + searchKey
    }
    if (right == 1) {
      searchKey =  searchKey + '*'
    }
    const res = await redisKeysApi({id: props.id, dbNo: selectDbNum.value, keyPattern: searchKey})
    const {list, max, total} = res.data
    keysRes.value.list = list
    keysRes.value.max = max
    keysRes.value.total = total
  } finally {
    dataFetchLoading.value = false
  }
}

const emit = defineEmits<{
  (e: 'close'): void
}>()


const showDashboardHandler = (show: boolean) => {
  emit('close')
}

const openUpdatePanelHandler = (key?: string) => {
  redisDataUpdatePanel.value.openHandler(key)
}

onMounted(() => {
  redisDbHandler()
})
</script>

<style lang="scss" scoped>

.redis-view {
  &-left {
    width: 150px;
    flex-shrink: 0;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    .title {
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .container {
      width: 100%;
      height: 100%;
    }
  }

  &-right {
    width: 100%;
    display: flex;
    flex-direction: column;

    .header {
      width: 100%;
      margin-top: 10px;
    }
  }
}


</style>