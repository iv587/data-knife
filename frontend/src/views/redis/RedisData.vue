<template>
    <div>
      <el-form size="mini" inline>
        <el-form-item label="实例:">
          <el-select @change="changeConnectionHandle" v-model="query.id">
            <el-option
              :value="item.id"
              :label="item.name"
              v-for="item in connectionList"
              :key="item.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="DB:">
          <el-select @change="changeDbNoHandle" v-model="query.dbNo">
            <el-option
              :label="`DB-${item}`"
              :value="item"
              v-for="item in dbList"
              :key="item"
            >
              <svg-icon icon-class="database" /> DB-{{item}}
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <el-form inline>
        <el-form-item>
          <el-input v-model="query.keyPattern">
            <template slot="append">
              <el-button type="success" @click="getData" icon="el-icon-search" />
            </template>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDlg()">添加</el-button>
        </el-form-item>
      </el-form>
      <section style="display: flex; justify-content: center">
        <el-table
          style="width: 80vw"
          v-loading="loading"
          :data="data.list"
          max-height="500">
          <el-table-column label="Key">
            <template slot-scope="scope">
              {{scope.row.key}}
            </template>
          </el-table-column>
          <el-table-column label="数据类型" align="center" width="80">
            <template slot-scope="scope">
              <el-tag :type="dataTagType[scope.row.type]" effect="plain">{{scope.row.type}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="TTL(s)" align="center" width="80" >
            <template slot-scope="scope">
              {{scope.row.ttl}}
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="150">
            <template slot-scope="scope">
              <el-button type="primary" icon="el-icon-edit" @click="openDlg(scope.row)" circle></el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="danger" icon="el-icon-delete" @click="deleteKeyHandler(scope.row)" circle></el-button>
            </template>
          </el-table-column>
        </el-table>
      </section>
      <el-dialog :visible.sync="dataDlg.show">
        <el-form v-model="dataDlg.data" :rules="dataDlg.rules" label-width="80px">
          <el-form-item label="Key:" prop="key">
            <el-row>
              <el-input  size="mini" style="width: 400px" :disabled="dataDlg.edit" v-model="dataDlg.data.key" />
            </el-row>
          </el-form-item>
          <el-form-item label="类型:">
            <el-select @change="dataTypeChange" size="mini" :disabled="dataDlg.edit" v-model="dataDlg.data.type">
              <el-option value="string" />
              <el-option value="list" />
              <el-option value="hash" />
              <el-option value="set" />
              <el-option value="zset" />
            </el-select>
          </el-form-item>
          <el-form-item label="TTL (秒):">
            <el-row :gutter="10">
              <el-col :span="5">
                <el-input size="mini" :disabled="!dataDlg.isUpdateTtl" placeholder="默认不过期" v-model="dataDlg.data.ttl">
                </el-input>
              </el-col>
              <el-col :span="6">
                <el-checkbox size="mini" v-if="!dataDlg.isUpdateTtl" v-model="dataDlg.isUpdateTtl">设置</el-checkbox>
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item v-if="dataDlg.data.type == 'string'" label="Value:">
             <el-input size="mini" type="textarea" v-model="dataDlg.data.value" ></el-input>
          </el-form-item>
          <el-form-item v-else label="Values:">
            <div style="width: 350px; padding-left: 5px" v-if="dataDlg.data.type === 'hash'">
              <el-row  :gutter="10">
                <el-col :span="8">
                  <el-input size="mini" v-model="dataDlg.addVal.field" placeholder="Key" />
                </el-col>
                <el-col :span="12">
                  <el-input size="mini" v-model="dataDlg.addVal.val" placeholder="value" />
                </el-col>
                <el-col :span="3">
                  <el-button @click="addVal" circle icon="el-icon-plus"  size="mini"/>
                </el-col>
              </el-row>
              <el-card  v-if="values.length > 0" :body-style="{ background:'#ecf0f1', paddingTop:'10px' }" shadow="never" style="max-height: 30vh; overflow-y: auto; margin-top: 10px">
                <el-row :gutter="10" v-for="item in values" :key="item.field">
                  <el-col :span="8"><el-input size="mini" v-model="item.field"></el-input></el-col>
                  <el-col :span="12"><el-input size="mini" v-model="item.val"></el-input></el-col>
                  <el-col :span="3"><el-button size="mini" icon="el-icon-minus" circle @click="removeVal(item.field)" /></el-col>
                </el-row>
              </el-card>
            </div>
            <div style="width: 250px; padding-left: 5px" v-if="['list','set'].indexOf(dataDlg.data.type) >= 0">
              <el-row  :gutter="10">
                <el-col :span="19">
                  <el-input size="mini" v-model="dataDlg.addVal.val" placeholder="value" />
                </el-col>
                <el-col :span="3">
                  <el-button @click="addVal" circle icon="el-icon-plus"  size="mini"/>
                </el-col>
              </el-row>
              <el-card  v-if="values.length > 0" :body-style="{ background:'#ecf0f1', paddingTop:'10px' }" shadow="never" style="max-height: 30vh; overflow-y: auto; margin-top: 10px">
                <el-row :gutter="10" v-for="item in values" :key="item.field">
                  <el-col :span="19"><el-input size="mini" v-model="item.val"></el-input></el-col>
                  <el-col :span="3"><el-button size="mini" icon="el-icon-minus" circle @click="removeVal(item.field)" /></el-col>
                </el-row>
              </el-card>
            </div>
            <div style="width: 250px; padding-left: 5px" v-if="dataDlg.data.type === 'zset'">
              <el-row  :gutter="10">
                <el-col :span="8">
                  <el-input size="mini" v-model="dataDlg.addVal.score" placeholder="score" />
                </el-col>
                <el-col :span="12">
                  <el-input size="mini" v-model="dataDlg.addVal.val" placeholder="member" />
                </el-col>
                <el-col :span="3">
                  <el-button @click="addVal" circle icon="el-icon-plus"  size="mini"/>
                </el-col>
              </el-row>
              <el-card  v-if="values.length > 0" :body-style="{ background:'#ecf0f1', paddingTop:'10px' }" shadow="never" style="max-height: 30vh; overflow-y: auto; margin-top: 10px">
                <el-row :gutter="10" v-for="item in values" :key="item.field">
                  <el-col :span="8"><el-input size="mini" v-model="item.score"></el-input></el-col>
                  <el-col :span="12"><el-input size="mini" v-model="item.val"></el-input></el-col>
                  <el-col :span="3"><el-button size="mini" icon="el-icon-minus" circle @click="removeVal(item.field)" /></el-col>
                </el-row>
              </el-card>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button @click="updateData" >保存</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
</template>

<script>
import { getRedisKeyList, getRedisKeyInfo, updateRedisData, getRedisServerList, getRedisDbs, deleteKey } from '@/api/redis'
import { mapActions, mapState } from 'vuex'
import { Message, MessageBox } from 'element-ui'

export default {
  name: 'RedisData',
  data() {
    return {
      dataTagType: {
        list: 'waring',
        set: 'info',
        hash: 'primary',
        zset: 'danger',
        string: 'success'
      },
      values: [],
      data: {
        list: []
      },
      dataDlg: {
        show: false,
        edit: false,
        isUpdateTtl: true,
        addVal: {
          field: '',
          val: '',
          score: ''
        },
        data: {
          values: []
        },
        rules: {
          key: [
            { required: true, message: '请输入key', trigger: 'blur' }
          ]
        }
      },
      connectionList: [],
      dbList: [],
      loading: false,
      query: {
        id: 0,
        dbNo: 0,
        keyPattern: ''
      }
    }
  },
  methods: {
    getData() {
      this.loading = true
      getRedisKeyList(
        this.query
      ).then(res => {
        this.data.list = res.data.list
        this.loading = false
      })
    },
    refreshData() {
      this.query.keyPattern = ''
      this.getData()
    },
    changeConnectionHandle() {
      this.currentServer({ id: this.query.id })
      this.setDbNo(0)
      this.query.dbNo = this.currDbNo
      this.refreshData()
    },
    changeDbNoHandle() {
      this.setDbNo(this.query.dbNo)
      this.refreshData()
    },
    // 更新数据
    updateData() {
      const _this = this
      this.dataDlg.data.values = JSON.stringify(this.values)
      this.dataDlg.data.isUpdateTtl = this.dataDlg.isUpdateTtl ? 1 : 0
      updateRedisData({
        ...this.dataDlg.data
      }).then(res => {
        Message.success({
          message: res.msg,
          type: 'success',
          duration: 1000,
          onClose: function() {
            _this.dataDlg.show = false
            _this.refreshData()
          }
        })
      })
    },
    // 打开添加或编辑对话框
    openDlg(data) {
      this.values = []
      this.dataDlg.data = { type: 'string', id: this.query.id, dbNo: this.query.dbNo }
      this.dataDlg.edit = false
      this.dataDlg.isUpdateTtl = true
      if (data) {
        getRedisKeyInfo({
          dbNo: this.query.dbNo,
          id: this.query.id,
          key: data.key
        }).then(res => {
          this.dataDlg.data = res.data
          this.dataDlg.data.id = this.query.id
          this.dataDlg.data.dbNo = this.query.dbNo
          this.dataDlg.edit = true
          this.dataDlg.isUpdateTtl = false
          this.dataDlg.show = true
          this.values = this.dataDlg.data.values
        })
      } else {
        this.dataDlg.show = true
      }
    },
    // 添加hash list set值
    addVal() {
      const data = this.dataDlg.addVal
      if (!data.val) {
        Message.error('请输入正确的值')
        return
      }
      // 判断key是否重复
      if (this.dataDlg.data.type === 'hash') {
        if (!data.field) {
          Message.error('请输入Key')
          return
        }
        for (const i in this.values) {
          if (this.values[i].field === data.field) {
            Message.error('Key重复')
            return
          }
        }
      }
      if (this.dataDlg.data.type === 'list') {
        data.field = new Date().getTime() + '-' + Math.random()
      }
      if (this.dataDlg.data.type === 'set') {
        for (const i in this.values) {
          if (this.values[i].val === data.val) {
            Message.error('不能有重复的值')
            return
          }
        }
        data.field = data.val
      }
      if (this.dataDlg.data.type === 'zset') {
        for (const i in this.values) {
          if (this.values[i].val === data.val) {
            Message.error('不能有重复的值')
            return
          }
        }
        data.field = data.val
      }
      this.values.unshift({
        ...data
      })
      this.dataDlg.addVal = {}
    },
    // 移除hash list set值
    removeVal(field) {
      let index = -1
      for (const i in this.values) {
        if (this.values[i].field === field) {
          index = i
          break
        }
      }
      if (index >= 0) {
        this.$nextTick(function() {
          this.values.splice(index, 1)
        })
      }
    },
    deleteKeyHandler(item) {
      const _this = this
      MessageBox.confirm('确认删除key=' + item.key + '数据？', {
        type: 'warning'
      }).then(res => {
        deleteKey({ id: _this.query.id, dbNo: _this.query.dbNo, key: item.key }).then(r => {
          Message.success({
            message: r.msg,
            type: 'success',
            duration: 1000,
            onClose: function() {
              _this.dataDlg.show = false
              _this.refreshData()
            }
          })
        })
      })
    },
    dataTypeChange() {
      this.values = []
      this.dataDlg.addVal = {
        field: '',
        val: '',
        score: ''
      }
    },
    ...mapActions(['currentServer', 'setMenuActiveIndex', 'setDbNo'])
  },
  mounted() {
    this.setMenuActiveIndex(this.$route.path)
    this.query.id = this.$store.state.redis.current.id
    getRedisServerList({}).then(res => {
      this.connectionList = res.data.list.map(function(item) {
        item.name = item.name ? item.name : item.addr
        return item
      })
    })
    this.query.dbNo = this.currDbNo
    getRedisDbs({ id: this.query.id }).then(res => {
      this.dbList = res.data.list
    })
    this.refreshData()
  },
  computed: {
    ... mapState({
      currDbNo: state => state.redis.current.dbNo
    })
  }
}
</script>

<style lang="scss">
</style>
