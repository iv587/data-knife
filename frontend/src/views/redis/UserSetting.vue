<template>
  <div>
    <el-row type="flex" justify="center">
      <el-col :span="8">
        <el-form ref="form" label-width="80px">
          <el-form-item label="用户名:">
            <el-input v-model="user.name" type="username" />
          </el-form-item>
          <el-form-item label="修改密码:">
            <el-input v-model="user.password" type="password" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="saveHandle">保存</el-button>
            <el-button type="danger" @click="logoutHandle">登出</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import { infoUser, updateUser } from '@/api/user'
import { Message } from 'element-ui'
export default {
  name: 'UserSetting',
  data() {
    return {
      user: {
        name: '',
        password: ''
      }
    }
  },
  mounted() {
    this.setMenuActiveIndex(this.$route.path)
    infoUser({}).then(res => {
      this.user = res.data
    })
  },
  methods: {
    logoutHandle() {
      this.removeToken().then(
        location.reload()
      )
    },
    saveHandle() {
      updateUser(this.user).then(res => {
        Message.success({
          message: res.msg,
          duration: 2000,
          onClose: function() {
            location.reload()
          }
        })
      })
    },
    ...mapActions(['removeToken', 'setMenuActiveIndex'])
  }
}
</script>

<style scoped>

</style>
