<template>
  <el-container class="logo" style="height: 100vh; background: #ecf0f1">
    <el-header height="230px" class="login-header" style="">
      <svg-icon class-name="icon" icon-class="redis" />
      <p class="title">Aha Redis</p>
    </el-header>
    <el-main style=" margin: 8vh auto">
      <el-form style="width: 30vw" label-width="120px">
        <el-form-item label="用户名">
          <span slot="label" style="color: grey">用户名:</span>
          <el-input v-model="param.name"  />
        </el-form-item>
        <el-form-item label="密码">
          <span slot="label" style="color: grey">密码:</span>
          <el-input type="password" v-model="param.password" />
        </el-form-item>
        <el-form-item>
          <el-button @click="loginHandle">登录</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import { login } from '@/api/user'
import { mapActions } from 'vuex'
export default {
  name: 'Login',
  data() {
    return {
      param: {
        name: '',
        password: ''
      }
    }
  },
  methods: {
    loginHandle () {
      let _this = this
      login(this.param).then(res => {
          this.setToken(res.data.token).then(r1 => {
            _this.$router.push({ path: '/'})
          })
      })
    },
    ...mapActions(['setToken'])
  }

}
</script>

<style scoped lang="scss">
  @import "~@/styles/myvariables.scss";
  .login-header {
    display: flex;
    justify-content: center;
    padding: 60px ;
    background: #1abc9c;
    .icon {
      $header-line: 5em;
      width: $header-line;
      height: $header-line;
      text-align: left;
    }
    .title {
      font-family: "Georgia", Tahoma, Sans-Serif;
      font-size: 60pt;
      color: #ffffff;
      vertical-align:middle;
      margin-left: 1vw;
    }
  }
</style>

