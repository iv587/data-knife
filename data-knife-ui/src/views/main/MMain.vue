<template>
  <el-container style="height: 100vh; width: 100vw; justify-content: space-between">
    <el-header class="dk-header">
      <dk-header/>
    </el-header>
    <el-main class="dk-main fill">
      <section class="fill" style="display: flex; flex-direction: column; justify-content: space-between;">
        <router-view :key="route.fullPath"/>
      </section>
      <section :class="{hidden: collapse}" class="menu-panel">
        <dk-menu @click="() => {menuStore.setCollapse(true)}" menu-cls="menu1" :menu-list="list"/>
      </section>
    </el-main>

    <!--    <el-container>
          <el-main class="fill" style="padding: 0;">
            <section class="m-main fill">
              <div class="m-main-view fill">
                <router-view :key="route.fullPath"/>
              </div>
              <section :class="{hidden: collapse}" class="menu-panel">
                <dk-menu :menu-list="list"/>
              </section>
            </section>
          </el-main>
        </el-container>-->

  </el-container>
  <!--
  <el-container style="height: 100vh; width: 100vw;">
    <el-header class="dk-header">
      <dk-header/>
    </el-header>
    <el-main style="padding: 0; height: 100%; width: 100%; background-color: var(--el-bg-color-page)">
      <div style="padding: 2px;height: 100%; width: 100%">
        <router-view :key="route.fullPath"/>
      </div>
    </el-main>
   <el-container style="height: 100%">
      <el-main style="padding: 10px; height: 100%; width: 100%; background-color: var(&#45;&#45;el-bg-color-page)">
        <div style="height: 100%; width: 100%">
          <router-view :key="route.fullPath"/>
        </div>
      </el-main>
      <section :class="{hidden: collapse}" class="menu-panel">
        <dk-menu :menu-list="list"/>
      </section>
    </el-container>
  </el-container>
  -->
</template>

<script lang="ts" setup>
import useMenuStore from "@/store/menu"
import {computed} from "vue";
import {useMenuData} from "@/views/main/menu";
import DkHeader from "@/components/layout/DkHeader.vue";

const menuStore = useMenuStore()

const {list, route} = useMenuData()

const collapse = computed(() => menuStore.collapse)

</script>

<style lang="scss" scoped>
.dk-main {
  margin: 0;
  padding: 0;
  height: calc(100vh - $headerHeight);
  width: 100vw;
  position: relative;
}

.m-main {
  position: relative;

  &-view {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    background-color: var(--el-bg-color-page);
    padding: 10px;
  }
}

.menu-panel {
  z-index: 9999;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  transition: left 200ms;
  background-color: var(--el-mask-color);
}

.hidden {
  left: -100%;
  transition: left 200ms;

}


.dk-header {
  flex-shrink: 0;
  height: $headerHeight;
  padding: 0;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--el-menu-border-color);
}
</style>