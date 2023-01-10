<template>
  <section>
    <section style="margin-bottom: 10px">
      <el-page-header @back="onClose">
        <template #title><h3>返回</h3></template>
        <template #content>
          <el-button size="small" circle :icon="Refresh" @click="redisInfoHandler"></el-button>
        </template>
      </el-page-header>
    </section>
    <section>
      <el-tabs v-loading="loading" v-if="infos.length > 0" :active-name="infos[0].label" type="border-card">
        <el-tab-pane v-for="section in infos" :key="section.label" :title="section.label" :label="section.label"
                     :name="section.label">
          <el-row type="flex" justify="start" v-for="item in section.items" :key="item.name">
            <el-col :span="10" style="color: #2980b9">{{ item.name }}:</el-col>
            <el-col :span="8" style="color: #3498db">{{ item.value }}</el-col>
          </el-row>
        </el-tab-pane>
      </el-tabs>
    </section>
  </section>
</template>

<script lang="ts" setup>
import {redisInfoApi} from '@/api/redis'
import {computed, onMounted, ref} from "vue";
import {RedisInfoSection} from "@/type";
import {Refresh } from '@element-plus/icons-vue'

const loading = ref(false)

const infos = ref<RedisInfoSection[]>([])

const props = defineProps<{
  dataId: number,
}>()

const emit = defineEmits<{
  (e: 'close'):void
}>()

const onClose = () => {
  emit('close')
}

const redisInfoHandler = async () => {
  loading.value = true
  const res = await redisInfoApi({id: props.dataId})
  infos.value = res.data
  loading.value = false
}

onMounted(() => {
  redisInfoHandler()
})

</script>

<style scoped>

</style>