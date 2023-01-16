<template>
  <el-dialog :width="width" title="更新" destroy-on-close v-model="open">
    <dk-form :form-data="formData" label-width="60px" @submit="updateHandle">
      <el-form-item prop="name" :rules="[{required: true, message: '请填写别名'}]" label="别名:">
        <el-input v-model="formData.name"/>
      </el-form-item>
      <el-form-item prop="addr" :rules="[{required: true, message: '请填写地址'}]" label="地址:">
        <el-input v-model="formData.addr"/>
      </el-form-item>
      <el-form-item prop="password" label="密码:">
        <el-input
          :disabled="formData.id != null && formData.id > 0 && (formData.upPass == null || formData.upPass == 0)"
          v-model="formData.password" type="password" show-password></el-input>
        <el-checkbox :true-label="1" :false-label="0" v-model="formData.upPass"
                     v-if="formData.id != null && formData.id > 0">修改密码
        </el-checkbox>
      </el-form-item>
      <el-form-item label="">
        <el-space>
          <el-button :loading="testStatus.loading" @click="testConnectionHandler">测试链接</el-button>
          <span v-if="testStatus.show" :style="`color: var(--el-color-${testStatus.type})`">{{ testStatus.msg }}</span>
        </el-space>
      </el-form-item>
    </dk-form>
  </el-dialog>
</template>

<script lang="ts" setup>
import DkForm from '@/components/DkForm.vue'
import {ref} from "vue";
import {getConnection, testConnection, updateConnection} from "@/api/connection";
import {Connection} from "@/type";
import message from "@/utils/message";

defineProps<{
  width?: string
}>()

const open = ref(false)

const formData = ref<{
  id?: number
  name?: string
  addr?: string
  password?: string,
  upPass?: number
}>({})

const testStatus = ref<{
  loading: boolean
  show: boolean
  type?: string,
  msg?: string
}>({
  show: false,
  loading: false
})

const emit = defineEmits<{
  (e: 'updateSuccess', id?: number) : void
}>()


const openHandler = async (id?: number) => {
  testStatus.value = {show: false, loading: false}
  formData.value = {}
  if (id != null && id > 0) {
    const res = await getConnection<Connection>({id,})
    formData.value.id = res.data.id
    formData.value.name = res.data.name
    formData.value.addr = res.data.addr
    formData.value.password = res.data.password
    formData.value.upPass = 0
  }
  open.value = true
}

const testConnectionHandler = async () => {
  testStatus.value.loading = true
  try {
    const res = await testConnection(formData.value)
    if (res.code == 1) {
      testStatus.value.type = 'success'
    } else if (res.code == 0) {
      testStatus.value.type = 'error'
    }
    testStatus.value.msg = res.msg
    testStatus.value.show = true
  } catch (err) {

  } finally {
    testStatus.value.loading = false
  }
}

const updateHandle = async () => {
  try {
    const res = await updateConnection(formData.value)
    await message.success(res.msg)
    open.value = false
    emit('updateSuccess', formData.value?.id)
  } catch (err) {
  }
}

defineExpose({
  openHandler
})


</script>

<style scoped>

</style>