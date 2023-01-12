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
          <el-button type="primary" @click="openConnectionUpdatePanelHandler()">添加</el-button>
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
            <el-button type="success" @click="openConnectionUpdatePanelHandler(row.id)" circle
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
import {computed, onMounted, ref} from "vue";
import {listConnection} from "@/api/connection";
import ConnectionUpdatePanel from "@/components/connection/ConnectionUpdatePanel.vue";
import {Connection} from "@/type";
import RedisCell from "@/components/redis/RedisCell.vue";
import {Delete, EditPen, InfoFilled} from "@element-plus/icons-vue";
import {useRouter} from "vue-router"
import useMenuStore from '@/store/menu'

const menuStore = useMenuStore()

const router = useRouter()

const connectionUpdatePanelRef = ref()

const list = ref<Connection[]>([])

const loading = ref(false)



const listConnectionHandler = async () => {
  loading.value = true
  try {
    const res = await listConnection()
    const data = res.data
    list.value = data.list
  } finally {
     loading.value = false
  }
}

const infoHandler = (conn: Connection) => {
  menuStore.updateMenuItem({
    id: conn.id + '',
    index: `/redis/${conn.id}`,
    title: conn.name,
    icon: 'redis-icon',
  })
  router.push({path: `/redis/${conn.id}`})
}

const onUpdateSuccessHandler = async (id?: number) => {
  await listConnectionHandler()
  if (id != null && id > 0) {
    for (const i in list.value) {
      if (list.value[i].id == id) {
        const conn = list.value[i]
        menuStore.updateMenuItem({
          id: conn.id + '',
          index: `/redis/${conn.id}`,
          title: conn.name,
          icon: 'redis-icon',
        }, true)
      }
    }
  }

}

const openConnectionUpdatePanelHandler = (id?: number) => {
  connectionUpdatePanelRef.value.openHandler(id)
}

onMounted(() => {
  listConnectionHandler()
})
</script>

<style lang="scss" scoped>
</style>