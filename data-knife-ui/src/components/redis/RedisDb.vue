<template>
  <section @click="onClick" :class="{selected: selected}" class="redis-bd">
    <el-icon>
      <redis-db-icon />
    </el-icon>
    <div>-DB</div>
    <p>{{dbNum}}</p>
  </section>
</template>

<script lang="ts" setup>
import RedisDbIcon from '@/icon/RedisDbIcon.vue'

const props = defineProps<{
  dbNum: number,
  selected: boolean,
}>()

const emit = defineEmits<{
  (e: 'selectDb', id: number):void
}>()

const onClick = () => {
  emit('selectDb', props.dbNum)
}

</script>

<style lang="scss" scoped>
.redis-bd {
  width: 100px;
  display: flex;
  align-items: center;
  font-size: 16px;
  font-width: 700;
  color: var(--el-color-info);
  margin-top: 15px;
  &:hover {
    color: var(--el-color-danger-light-3);
    cursor: pointer;
    transform: scaleY(1.18);
  }
}
.selected {
  color: var(--el-color-danger);
}
</style>