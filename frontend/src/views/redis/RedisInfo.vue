<template>
    <div>
      <el-tabs :active-name="info[0].section" type="border-card" >
        <el-tab-pane v-for="item in info" :key="item.section" :title="item.section" :label="item.section" :name="item.section">
          <el-row type="flex" justify="start" v-for="item1 in item.values" :key="item1.feild">
            <el-col :span="10" style="color: #2980b9">{{item1.field}}:</el-col>
            <el-col :span="8" style="color: #3498db">{{item1.val}}</el-col>
          </el-row>
        </el-tab-pane>
      </el-tabs>
      <!--<el-card style="margin-bottom: 10px" v-for="item in info" :key="item.section" :title="item.section">
        <template slot="header">
          <span style="color: #34495e;">{{item.section}}</span>
        </template>
        <el-row type="flex" justify="start" v-for="item1 in item.values" :key="item1.feild">
          <el-col :span="10" style="color: #2980b9">{{item1.field}}:</el-col>
          <el-col :span="8" style="color: #3498db">{{item1.val}}</el-col>
        </el-row>
      </el-card>-->
    </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import { redisInfo } from '@/api/redis'
export default {
  name: 'RedisInfo',
  data() {
    return {
      info: ''
    }
  },
  computed: {
    ... mapState ({
      id: state => state.redis.current.id,
      dbNo: state => state.redis.current.dbNo
    })
  },
  methods: {
    ... mapActions([ 'setMenuActiveIndex' ])
  },
  mounted() {
    this.setMenuActiveIndex(this.$route.path)
    redisInfo({
      id: this.id,
      dbNo: this.dbNo
    }).then( res => {
      this.info = res.data
    })
  }
}
</script>

<style scoped>

</style>
