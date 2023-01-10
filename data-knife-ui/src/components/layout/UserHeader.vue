<template>
<section style="height: 100%; display: flex; align-items: center; justify-content: center">
  <el-dropdown @command="onCommand">
    <el-space style="cursor: pointer">
      <span style="color: var(--el-color-info);">{{userInfo.userName}}</span>
      <el-avatar :icon="UserFilled"></el-avatar>
    </el-space>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="setting"><el-icon><Setting /></el-icon>设置密码</el-dropdown-item>
        <el-dropdown-item command="logout"><el-icon><SwitchButton /></el-icon>登出系统</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
  <el-dialog width="35vw" title="修改密码" destroy-on-close v-model="showSettingDlg">
    <el-form :model="formData"  ref="settingForm" label-width="auto">
      <el-form-item prop="oldPassword" label="旧密码:" :rules="[{required: true, message: '请输入旧密码'}]">
        <el-input v-model="formData.oldPassword" type="password" show-password placeholder="请输入旧密码"></el-input>
      </el-form-item>
      <el-form-item prop="newPassword" label="新密码:" :rules="[{required: true, message: '请输入新密码'}]">
        <el-input  v-model="formData.newPassword" type="password" show-password placeholder="请输入新密码"></el-input>
      </el-form-item>
      <el-form-item prop="rePassword" label="确认密码:" :rules="[{required: true, validator: validatePass2}]">
        <el-input v-model="formData.rePassword" type="password" show-password placeholder="请再次输入新密码"></el-input>
      </el-form-item>
      <el-form-item label=" ">
        <el-button>取消</el-button>
        <el-button @click="setPasswordHandler" type="primary">确定</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</section>
</template>

<script lang="ts" setup>
import {UserFilled, Setting, SwitchButton } from '@element-plus/icons-vue'
import useUserStore from "@/store/user";
import {computed, ref} from "vue";
import {ElMessageBox, FormInstance} from "element-plus";
import message from "@/utils/message";
import router from "@/router";
import {updatePasswordApi} from "@/api/user";

const userStore = useUserStore()

const settingForm = ref<FormInstance>()

const formData = ref<{
  oldPassword: string
  newPassword: string
  rePassword: string
}>({
  oldPassword: '',
  newPassword: '',
  rePassword: ''
})

const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入新密码'))
  } else if (value !== formData.value.newPassword) {
    callback(new Error("两次密码不一致!"))
  } else {
    callback()
  }
}

const setPasswordHandler = async () => {
  try {
    await settingForm.value?.validate()
    await ElMessageBox.confirm('将要设置密码', '提示', {type: "warning"})
    const res = await updatePasswordApi(formData.value)
    await message.success(res.msg)
    showSettingDlg.value = false
  }catch (e) {
    
  }
}

const userInfo = computed(() => {
  return {
    userName:   userStore.userName
  }
})

const showSettingDlg = ref(false)

const onCommand = async (command: string|number|object) => {
  if (command === 'setting') {
    showSettingDlg.value = true
  }
  if (command === 'logout') {
    await ElMessageBox.confirm('是否退出系统', '提示', {type: "warning"})
    userStore.logout()
    window.location.reload()
  }
}

</script>

<style scoped>

</style>