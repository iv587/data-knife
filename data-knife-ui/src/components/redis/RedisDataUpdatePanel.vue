<template>
  <el-dialog destroy-on-close v-model="show">
    <template #header>
      <el-space v-if="formData.key != null && formData.key != ''">
        <el-link type="info">
          正在更新:
        </el-link>
        <el-link type="danger">{{formData.key}}</el-link>
      </el-space>
      <el-link type="info" v-else>添加数据</el-link>
    </template>
    <dk-form v-loading="loading" :form-data="formData" label-width="80px" @submit="updateDataHandler">
      <el-form-item label="Key:" prop="key" :rules="[{required: true, message: '请填写key'}]">
        <el-input :disabled="editMode" v-model="formData.key" placeholder="设置缓存key" clearable/>
      </el-form-item>
      <el-form-item label="类型:" prop="type" :rules="[{required: true, message: '请选择type'}]">
        <el-select :disabled="editMode" v-model="formData.type" placeholder="设置数据类型">
          <el-option v-for="item in keyTypes" :value="item" :key="item"/>
        </el-select>
      </el-form-item>
      <el-form-item label="TTL (秒):">
        <el-space>
          <el-input-number :disabled="editMode && formData.isUpdateTtl == 0" v-model="formData.ttl" :precision="0"
                           :min="0" controls-position="right" placeholder="0表示永不过期"/>
          <el-checkbox v-if="editMode" v-model="formData.isUpdateTtl" :true-label="1" :false-label="0">设置
          </el-checkbox>
        </el-space>
      </el-form-item>
      <el-form-item v-if="formData.type == 'string'" label="Value:">
        <el-input type="textarea" placeholder="设置数据" v-model="formData.value"></el-input>
      </el-form-item>
      <el-form-item v-if="formData.type == 'hash'" label="values">
        <redis-hash-update ref="redisHashUpdate" :values="values"/>
      </el-form-item>
    </dk-form>
  </el-dialog>
</template>

<script lang="ts" setup>

import {ref} from "vue";
import DkForm from "@/components/DkForm.vue";
import {redisKeyInfoApi, redisKeyUpdateApi} from "@/api/redis"
import message from "@/utils/message";
import {RedisKvPair} from "@/type";
import RedisHashUpdate from "@/components/redis/update/RedisHashUpdate.vue";

const show = ref(false)

const loading = ref(false)

const keyTypes = ref([
  'string',
  'list',
  'hash',
  'set',
  'zset'
])

const props = defineProps<{
  connectionId: number,
  dbNo: number,
}>()

const formData = ref<{
  id: number,
  dbNo: number,
  key: string,
  type: string,
  ttl: number,
  value?: string,
  isUpdateTtl: number,
}>({
  id: props.connectionId,
  dbNo: props.dbNo,
  key: '',
  type: '',
  ttl: 0,
  isUpdateTtl: 1,
})

const values = ref<RedisKvPair[]>([])

const editMode = ref(false)

const redisHashUpdate = ref()

const emit = defineEmits<{
  (e: 'updateSuccess'): void
}>()

const openHandler = async (key?: string) => {
  editMode.value = false
  formData.value.key = ''
  formData.value.type = 'string'
  formData.value.ttl = 0
  formData.value.isUpdateTtl = 1
  values.value = []
  if (key && key != '') {
    try {
      loading.value = true
      show.value = true
      const res = await redisKeyInfoApi({id: props.connectionId, dbNo: props.dbNo, key: key})
      const data = res.data
      formData.value.key = data.key
      formData.value.type = data.type
      formData.value.ttl = data.ttl
      formData.value.value = data.value
      formData.value.isUpdateTtl = 0
      editMode.value = true
      values.value = data.values
    } finally {
      loading.value = false
    }
  } else {
    show.value = true
  }
}


const updateDataHandler = async () => {
  try {
    let postData: any = {...formData.value}
    loading.value = true
    if (formData.value.type === 'hash') {
      const {updateList, deleteList} =  redisHashUpdate.value.getActualUpdateAndDelList()
      postData.values = JSON.stringify(updateList)
      postData.delValues = JSON.stringify(deleteList)
    }
    const res = await redisKeyUpdateApi(postData)
    await message.success(res.msg)
    show.value = false
    emit('updateSuccess')
  } finally {
    loading.value = false
  }
}

defineExpose({
  openHandler
})

</script>

<style scoped>

</style>