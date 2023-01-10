<template>
  <section class="redis-cell">
    <div class="icon">
      <el-icon>
        <redis-icon />
      </el-icon>
    </div>
    <div class="content">
      <h3 class="name">{{data.name}}</h3>
      <p class="addr">{{data.addr}}</p>
    </div>
  </section>
</template>

<script lang="ts" setup>
import RedisIcon from '@/icon/RedisIcon.vue'
import {Connection} from "@/type";

defineProps<{
  data: Connection
}>()

</script>

<style lang="scss" scoped>
.redis-cell {
  display: flex;
  .icon {
    font-size: 60px;
    display: flex;
    justify-content: center;
    justify-items: center;
    margin-right: 5px;
  }
  .content {
    display: flex;
    flex-direction: column;
    justify-content: center;
    flex-grow: 1;
    .addr {
      color: var(--el-color-info);
    }
  }
}
</style>