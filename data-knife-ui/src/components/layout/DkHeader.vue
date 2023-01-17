<template>

  <section style="padding: 0 20px 0 10px; display: flex; justify-content: space-between; align-content: center; flex: 1">
    <div style="display: flex; align-items: center">
      <icon-button v-if="collapse" :size="30" @click="onToggleCollapse">
        <Expand  />
      </icon-button>
      <icon-button v-else :size="30" @click="onToggleCollapse">
        <Fold />
      </icon-button>
      <el-divider direction="vertical" />
      <el-space v-if="activeMenuItem">
        <el-icon v-if="activeMenuItem.icon">
          <component :is="activeMenuItem.icon"/>
        </el-icon>
        <span>{{activeMenuItem.title}}</span>
      </el-space>
    </div>
    <user-header  />
  </section>
<!--  <el-header class="dk-header">

  </el-header>-->
</template>

<script lang="ts" setup>
import UserHeader from '@/components/layout/UserHeader.vue'
import {Expand, Fold } from '@element-plus/icons-vue'
import IconButton from "@/components/IconButton.vue";
import MenuStore from '@/store/menu'
import {computed} from "vue";

const menuStore = MenuStore()
const onToggleCollapse = () => menuStore.toggleCollapse()
const collapse = computed(() =>{return menuStore.collapse} );
const activeMenuItem = computed(() => {
  return menuStore.getActiveMenuItem

})
</script>

<style lang="scss" scoped>
.dk-header {
  flex-shrink: 0;
  height: $headerHeight;
  padding: 0;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--el-menu-border-color);
}
</style>