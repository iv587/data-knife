<template>
<el-form status-icon ref="form" :model="formData" :label-width="labelWidth">
  <slot></slot>
  <el-form-item>
    <el-button @click="onReset">取消</el-button>
    <el-button type="primary" @click="submitForm()">确定</el-button>
  </el-form-item>
</el-form>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {FormInstance} from "element-plus";

defineProps<{
  labelWidth?: string,
  formData?: any
}>()

const emit = defineEmits<{
  (e: 'submit') :void
}>()

const form = ref<FormInstance>()


const onReset = () => {
  form.value?.resetFields()
}

const submitForm = async () => {
  try {
    await form.value?.validate()
    emit('submit')
  }catch (err) {

  }

}


</script>

<style scoped>

</style>