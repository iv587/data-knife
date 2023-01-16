<template>
  <page-container>
    <section v-if="showDashBoard" style="height: 100%; width: 100%">
      <redis-info-view  :data-id="dataId" @close="showDashboardHandler(false)"/>
    </section>
    <m-redis-data-view v-show="!showDashBoard"  :id="dataId" @close="showDashboardHandler(true)"/>
  </page-container>
</template>

<script lang="ts" setup>
import PageContainer from "@/components/layout/PageContainer.vue";
import {ref} from "vue";
import {useRouter} from 'vue-router'
import RedisInfoView from "@/views/redis/RedisInfoView.vue";
import MRedisDataView from "@/views/redis/MRedisDataView.vue";
const router = useRouter()
const id = router.currentRoute.value.params.id
const dataId = ref(parseInt(id as string))

const showDashBoard = ref(false)

const showDashboardHandler = (show: boolean) => {
  showDashBoard.value = show
}


</script>

<style lang="scss" scoped>

.redis-view {
  &-left {
    width: 150px;
    flex-shrink: 0;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    .title {
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .container {
      width: 100%;
      height: 100%;
    }
  }

  &-right {
    width: 100%;
    display: flex;
    flex-direction: column;

    .header {
      width: 100%;
      margin-top: 10px;
    }
  }
}


</style>