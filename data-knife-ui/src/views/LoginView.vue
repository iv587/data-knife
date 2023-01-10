<template>
  <section  style="display: flex; flex-direction: column; align-items: center ;width: 100vw; height: 100vh;">
    <section class="header">
      <p class="title">Data Knife</p>
    </section>
    <section v-if="hadInit" class="login">
      <el-form  ref="loginFormRef" :model="loginForm">
        <el-form-item prop="userName" :rules="[{required: true, message: '请输入登录名'}]">
          <el-input placeholder="请输入登录名" v-model="loginForm.userName">
            <template #prefix>
              <el-icon><User /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :rules="[{required: true, message: '请输入密码'}]">
          <el-input placeholder="请输入密码" v-model="loginForm.password" type="password" show-password >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 100%" type="primary" @click="loginHandler">登录</el-button>
        </el-form-item>
      </el-form>
      <el-alert type="error" :closable="false" title="登录Data Knife" description="一款数据操作利器"></el-alert>
    </section>
    <section class="app-init" v-else>
      <p class="title">初次使用，请先创建用户~</p>
      <el-form  ref="initFormRef" :model="initForm">
        <el-form-item prop="userName" :rules="[{required: true, message: '请填写用户名'}]">
          <el-input placeholder="请填写用户名" v-model="initForm.userName">
            <template #prefix>
              <el-icon><User /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password" :rules="[{required: true, message: '请设置密码'}]">
          <el-input placeholder="请设置密码" v-model="initForm.password" type="password" show-password >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="rePassword" :rules="[{required: true, validator: validatePass2}]">
          <el-input placeholder="请再次输入密码" v-model="initForm.rePassword" type="password" show-password >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 100%" type="primary" @click="appInitHandler">创建</el-button>
        </el-form-item>
      </el-form>
    </section>
  </section>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {appCheckApi, appInitApi, loginApi} from "@/api/user";
import {FormInstance} from "element-plus";
import useUserStore from '@/store/user'
import {Lock, User} from '@element-plus/icons-vue'
import {useRouter} from "vue-router";
import message from "@/utils/message";

const loginFormRef = ref<FormInstance>()

const userStore = useUserStore()

const router = useRouter()

const loginForm = ref<{
  userName: string,
  password: string,
}>({
  userName: '',
  password: ''
})

const initForm = ref<{
  userName: string
  password: string
  rePassword: string
}>({
  userName: '',
  password: '',
  rePassword: ''
})

const initFormRef = ref<FormInstance>()

const hadInit = ref(true)

const loginHandler = async () => {
  try {
    await loginFormRef.value?.validate()
    const res = await loginApi(loginForm.value)
    userStore.setLoginRes(res.data)
    await router.push({path: '/'})
  }catch (err) {
  }
}

const appCheckHandler = async () => {
  const res = await appCheckApi({})
  const {init} = res.data
  hadInit.value = init > 0 ? true : false
}

const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入新密码'))
  } else if (value !== initForm.value.password) {
    callback(new Error("两次密码不一致!"))
  } else {
    callback()
  }
}

const appInitHandler = async () => {
  await initFormRef.value?.validate()
  await appInitApi(initForm.value)
  await message.success('初始化成功')
  window.location.reload()
}

onMounted(async () => {
  await appCheckHandler()
})


</script>

<style lang="scss" scoped>
.header {
  width: 100%;
  background-color: #1abc9c;
  display: flex;
  justify-content: center;
  padding: 60px ;
  .title {
    font-family: "Georgia", Tahoma, Sans-Serif;
    font-size: 60pt;
    color: #ffffff;
    vertical-align:middle;
    margin-left: 1vw;
  }
}
.login {
  margin-top: 10vh;
  width: 20vw;
}
.app-init {
  margin-top: 10vh;
  width: 20vw;
  .title {
    color: var(--el-color-primary);
    font-size: 20px;
    margin-bottom: 10px;
  }
}
</style>