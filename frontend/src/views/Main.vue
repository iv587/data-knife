<template>
  <el-container style="height: 100vh; background: #ecf0f1">
    <el-aside width="250px" >
      <el-menu
        style="overflow: auto; height: 100vh"
        background-color="#2f3640"
        text-color="#ffffff"
        :default-active="activeIndex"
        router
      >
        <div class="logo">
           <svg-icon class-name="icon" icon-class="redis" ></svg-icon>&nbsp;
           <p class="text">Aha Redis</p>
          <p></p>
        </div>
        <el-menu-item index="/redis/connections" >连接管理</el-menu-item>
        <el-submenu v-if="serverId" index="2">
          <template slot="title">{{ this.name }}</template>
          <el-menu-item index="/redis/info">概览</el-menu-item>
          <el-menu-item index="/redis/db">数据</el-menu-item>
        </el-submenu>
        <el-menu-item index="/user/setting">用户设置</el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header height="50px" class="navbar">
        <div class="right-menu">
          <svg-icon class-name="user" icon-class="user"></svg-icon>
        </div>
      </el-header>
      <el-main style="padding: 2px 5px">
        <div style="padding: 10px ; background: #ffffff;">
          <router-view />
        </div>
        <div class="footer">
          <el-link href type="info">
            <svg-icon class-name="githubicon" icon-class="github" ></svg-icon><span class="title">Git Hub</span>
          </el-link>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'Main',
  data() {
    return {
    }
  },
  mounted() {
    this.$router.push({ path: this.activeIndex })
  },
  computed: {
    ... mapState({
      activeIndex: state => state.menu.activeIndex,
      serverId: state => state.redis.current.id,
      name: state => state.redis.current.name,
    })
  }
}
</script>

<style lang="scss" scoped>
  @import "~@/styles/myvariables.scss";
  .logo {
    margin-top: 2vh;
    margin-bottom: 2vh;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    .icon {
      $iconSize: 3em;
      width: $iconSize;
      height: $iconSize;
      margin-right: 1vw;
     }
     .text {
       color: #1abc9c;
       font-family: "Georgia", Tahoma, Sans-Serif;
       font-size: 20pt;
     }
  }
  .navbar {
    overflow: hidden;
    position: relative;
    background: #fff;
    box-shadow: 0 1px 4px rgba(0,21,41,.08);
    .right-menu{
      float: right;
      height: 100%;
      line-height: 50px;
      display: flex ;
      align-items: center;
      .user {
        width: 2em;
        height: 2em;
      }
    }
  }
  .footer {
    display: flex;
    justify-content: center;
    margin-top: 10vh ;
    align-items: center;
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    .githubicon {
        width: 1.7em;
        height: 1.7em;
      }
    .title {
      font-size: 15pt;
    }
  }


</style>
