<template>
  <div>
    <el-form>
      <el-form-item><el-button @click="openDlg()">创建</el-button></el-form-item>
    </el-form>
    <el-table
        :data="instances.list"
        highlight-current-row
        :show-header="false"
      >
        <el-table-column width="60">
          <template >
            <svg-icon class-name="table-redis-icon" icon-class="redis" />
          </template>
        </el-table-column>
        <el-table-column style="border: none" align="left">
          <template slot-scope="scope">
            <p style="color: #606266; font-size: 15pt;">{{scope.row.name}}</p>
            <p style="color: #909399">{{scope.row.addr}}</p>
          </template>
        </el-table-column>
        <el-table-column align="center" width="250">
          <template slot-scope="scope">
            <el-button @click="rowClick(scope.row)" type="primary" icon="el-icon-info" circle></el-button>
            <el-divider direction="vertical" />
            <el-button @click="openDlg(scope.row)" type="success" icon="el-icon-edit" circle></el-button>
            <el-divider direction="vertical" />
            <el-button
              @click="deleteData(scope.row)"
              type="danger"
              icon="el-icon-delete"
              circle></el-button>
          </template>
        </el-table-column>
      </el-table>
  <el-dialog
    :title= "editDlg.title"
    :visible.sync="editDlg.show"
    width="40%">
    <el-form ref="form" :model="editDlg.data" label-width="80px">
      <el-form-item label="别名:">
        <el-input v-model="editDlg.data.name"></el-input>
      </el-form-item>
      <el-form-item label="地址:">
        <el-input v-model="editDlg.data.addr"></el-input>
      </el-form-item>
      <el-form-item v-if="editDlg.isEdit">
        <el-checkbox  :true-label="1" :false-label="0" v-model="editDlg.upPass" >修改密码</el-checkbox>
      </el-form-item>
      <el-form-item v-if="!editDlg.isEdit || editDlg.upPass" label="密码">
        <el-input type="password" v-model="editDlg.data.password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button v-if="!editDlg.isEdit || editDlg.upPass" size="mini" @click="testConnHandle">测试连接</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="updateConnection">确 定</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
  </div>
</template>

<script>
  import {getRedisServerList, connectionUpdate, testConnection, deleteConnection} from '@/api/redis'
import { mapActions } from 'vuex'
import { Message, MessageBox } from 'element-ui'

export default {
  name: 'RedisServer',
  data() {
    return {
      instances: {
        list: []
      },
      editDlg: {
        title: '',
        isEdit: false,
        show: false,
        upPass: 0,
        data: {
          id: 0,
          password: '',
          name: '',
          addr: ''
        }
      }
    }
  },
  methods: {
    getRedisServer() {
      getRedisServerList().then(res => {
        this.instances = res.data
      })
    },
    openDlg(item) {
      this.editDlg.isEdit = false
      this.editDlg.upPass = 0 ;
      this.editDlg.data = {
        id: 0,
        password: '',
        name: '',
        addr: ''
      }
      if (item) {
        this.editDlg.title = '编辑' + item.name
        this.editDlg.data.id = item.id ;
        this.editDlg.data.addr = item.addr ;
        this.editDlg.data.name = item.name ;
        this.editDlg.isEdit = true
      }else {
        this.editDlg.title = '添加'
      }
      this.editDlg.show = true
    },
    updateConnection() {
      const upPass = this.editDlg.upPass
      const data = this.editDlg.data
      let _this = this
      connectionUpdate({
         upPass,...data
      }).then(res => {
        Message.success({
          message: res.msg,
          type: 'success',
          duration: 1000,
          onClose: function () {
            _this.editDlg.show = false
            _this.getRedisServer()
          }
        })
      })
    },
    rowClick(row) {
      this.currentServer({ id: row.id, dbNo: 0, name: row.name})
      this.$router.push({ path: '/redis/info', params: {}})
    },
    deleteData(item) {
      let _this = this
      MessageBox.confirm('确认删除连接'+item.name+'数据？',{
        type: 'warning'
      }).then(res => {
        deleteConnection({id: item.id }).then(r => {
          Message.success({
            message: r.msg,
            type: 'success',
            duration: 1000,
            onClose: function () {
              _this.currentServer1(item.id)
              _this.getRedisServer()
            }
          })
        })
      })
    },
    testConnHandle() {
      testConnection(this.editDlg.data).then(res => {
        Message.success({
          message: res.msg,
          type: 'success',
          duration: 1000,
          onClose: function () {
            _this.editDlg.show = false
            _this.getRedisServer()
          }
        })
      })
    },
    ...mapActions(['currentServer', 'currentServer1', 'setMenuActiveIndex'])
  },
  mounted() {
    this.setMenuActiveIndex(this.$route.path)
    this.getRedisServer()
  },
  computed: {
    serverId() {
      console.log(this.$store.state.redis.current.serverId)
      return this.$store.state.redis.current.serverId
    }
  }
}
</script>

<style scoped>
  .table-redis-icon {
    width: 2.5em;
    height: 2.5em;
  }
</style>
