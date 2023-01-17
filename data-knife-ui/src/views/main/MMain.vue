<template>
  <el-container style="height: 100vh; width: 100vw; justify-content: space-between">
    <el-header class="dk-header">
      <dk-header/>
    </el-header>
    <el-main class="dk-main fill">
      <section class="fill" style="display: flex; flex-direction: column; justify-content: space-between;">
        <router-view v-slot="{Component, route}">
          <transition name="fade" mode="out-in">
            <keep-alive>
              <component :key="route.path" :is="Component"></component>
            </keep-alive>
          </transition>
        </router-view>
      </section>
      <section :class="{hidden: collapse}" class="menu-panel">
        <dk-menu @click="() => {menuStore.setCollapse(true)}" menu-cls="menu1" :menu-list="list"/>
      </section>
    </el-main>
  </el-container>
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


.fade-enter-active,
.fade-leave-active {
  transition: opacity 100ms ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>