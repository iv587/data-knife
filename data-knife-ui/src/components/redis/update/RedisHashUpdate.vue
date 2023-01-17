<template>
  <div style="width: 100%">
    <el-row :gutter="10">
      <el-col :span="11">
        <el-input v-model="addData.field" placeholder="field"/>
      </el-col>
      <el-col :span="11">
        <el-input v-model="addData.val" placeholder="value"/>
      </el-col>
      <el-col :span="2">
        <el-button @click="onAdd(addData.field, addData.val)" circle :icon="Plus" size="small"/>
      </el-col>
    </el-row>
  </div>
  <div style="margin-top: 5px ;width: 100%;  border: 1px solid var(--el-border-color)">
    <section style="display: flex; align-items: center ;padding: 5px; background-color: var(--el-bg-color-page);">
      <el-input style="width: 50%" v-model="searchKey" placeholder="输入field进行搜索" @input="onSearch"/>
      <p style="margin-left: 10px; color: var(--el-color-success)">共有{{ values.length }}个字段</p>
    </section>
    <el-table
      @selection-change="onSelectChange"
      size="small" :data="tableList" :max-height="mobileDevice ? '20vh':'35vh'">
      <el-table-column prop="field" label="filed">
      </el-table-column>
      <el-table-column label="value">
        <template #default="{row}">
          <div class="edit-cell" v-if="editFiled == row.field">
            <el-input v-model="editData.val"></el-input>
            <el-space class="edit-btn">
              <el-link @click="onAdd(editData.field, editData.val)" :underline="false" type="success"
                       :icon="Select"></el-link>
              <el-link @click="cancelEdit()" :underline="false" type="danger" :icon="CloseBold"></el-link>
            </el-space>
          </div>
          <p v-else>{{ row.val }}</p>
        </template>
      </el-table-column>
      <el-table-column width="180" align="center">
        <template #default="{row}">
          <el-button plain type="success" @click="onEditHandler(row.field, row.val)" circle :icon="EditPen"
                     size="small"/>
          <el-button plain type="danger" @click="onRemove(row.field)" circle :icon="Delete" size="small"/>
        </template>
      </el-table-column>
    </el-table>
  </div>

</template>

<script lang="ts" setup>

import {RedisKvPair} from "@/type";
import {CloseBold, Delete, EditPen, Plus, Select} from "@element-plus/icons-vue";
import {h, onMounted, ref} from "vue";
import message from "@/utils/message";
import {ElMessageBox} from "element-plus";
import RedisHashConfirmTips from "@/components/redis/update/RedisHashConfirmTips.vue";
import {isMobile} from "@/utils/http";

const mobileDevice = isMobile()

const props = defineProps<{
  values: RedisKvPair[]
}>()

const addData = ref<RedisKvPair>({
  field: '',
  val: '',
  score: '',
})

const editData = ref<{
  field: string
  val: string,
  oldVal: string
}>({
  field: '',
  val: '',
  oldVal: ''
})

// 汇总list
const dataList = ref<RedisKvPair[]>([])
dataList.value = props.values.map(it => {
  return {
    field: it.field,
    score: it.score,
    val: it.val
  }
})

// 渲染table的list
const tableList = ref<RedisKvPair[]>([])

const editFiled = ref<string>('')

const onEditHandler = (field: string, val: string) => {
  editFiled.value = field
  editData.value = {
    field: field,
    val: val,
    oldVal: val
  }
}

const selectIds = ref<string[]>([])

const cancelEdit = () => {
  editFiled.value = ''
}
// 需要更新的数据列表
const updateMap = ref<Map<string, RedisKvPair>>(new Map<string, RedisKvPair>())
// 需要删除的数据列表
const deleteSet = ref<Set<string>>(new Set<string>())

// 添加
const onAdd = async (field: string, val: string) => {
  if (field == '') {
    await message.error('请填写field')
    return
  } else if (val == '') {
    await message.error('请填写val')
    return
  }
  try {
    const index = dataList.value.findIndex(item => {
      return item.field === field
    })
    const addItem: RedisKvPair = {
      field, val, score: ''
    }
    if (index < 0) {
      await ElMessageBox.confirm(h(RedisHashConfirmTips, {
        field: field,
        val: val
      }, ''), '将要添加数据')
    } else {
      const item = dataList.value[index]
      if (props.values.find(it => it.field === field)) {
        await ElMessageBox.confirm(h(RedisHashConfirmTips, {
          field: field,
          val: val,
          oVal: item.val
        }, '更新确认'), '将执行修改操作')
      }
      dataList.value.splice(index, 1)
      editFiled.value = ''
    }
    dataList.value.unshift(addItem)
    addData.value = {field: '', val: '', score: ''}
    updateMap.value.set(addItem.field, addItem)
    deleteSet.value.delete(addItem.field)
    onSearch()
  } catch (err) {
  }
}

const onSelectChange = (selection: RedisKvPair[]) => {
  selectIds.value = selection.map(it => it.field)
  console.log(selectIds.value)
}

// 删除
const onRemove = async (field?: string) => {
  let deleteIds: string[] = []
  if (field != null) {
    deleteIds.push(field)
  }
  try {
    await ElMessageBox.confirm('是否删除以下Field: ' + deleteIds, '删除确认', {type: "warning"})
    for (const i in deleteIds) {
      deleteSet.value.add(deleteIds[i])
      updateMap.value.delete(deleteIds[i])
      dataList.value.forEach((value, index, array) => {
        if (value.field == deleteIds[i]) {
          array.splice(index, 1)
        }
      })
    }
  } catch (err) {
  }
  onSearch()
}

const searchKey = ref<string>('')

const onSearch = () => {
  if (searchKey.value == '') {
    tableList.value = dataList.value
  }
  tableList.value = dataList.value.filter(it => {
    return it.field.includes(searchKey.value)
  })
}
onMounted(() => {
  onSearch()
})

const getActualUpdateAndDelList = (): { updateList: RedisKvPair[], deleteList: string[] } => {
  return {
    updateList: Array.from(updateMap.value.values()), deleteList: Array.from(deleteSet.value.values())
  }
}

defineExpose({
  onSearch,
  getActualUpdateAndDelList
})

</script>

<style lang="scss" scoped>
.edit-cell {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .edit-btn {
    margin-left: 10px;
    flex-shrink: 0;
  }
}
</style>