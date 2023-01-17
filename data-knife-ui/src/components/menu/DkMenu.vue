<template>
  <el-menu class="menu"
           router
           @select="onClick"
           :collapse="collapse"
           style="height: 100%"
           :default-active="activeIndex">
    <menu-logo />
    <dk-menu-item v-for="(item, index) in menuList" :menu-item="item" :key="index" />
  </el-menu>
</template>

<script lang="ts" setup>
import DkMenuItem from '@/components/menu/DkMenuItem.vue'
import {computed} from 'vue'
import {MenuItem} from "@/type";
import useMenuStore from '@/store/menu'
import {useRouter} from "vue-router";
import {isMobile} from "@/utils/http";

const menuStore = useMenuStore()
const router = useRouter()

const emit = defineEmits<{
  (e: 'click'):void
}>()

defineProps<{menuList: MenuItem[],
menuCls?: string
}>()

const activeIndex = computed(() => {
  return menuStore.activeIndex
})

const mobileDevice = isMobile()

const collapse= computed(() => {
  return mobileDevice ? null : menuStore.collapse
})

const onClick = (index: string) => {
  menuStore.setActiveIndex(index)
  emit('click')
}

</script>

<style lang="scss">
@media screen and (min-width: 700px){
  .menu:not(.el-menu--collapse) {
    width: 250px;
  }
}

@media screen and (max-width: 700px){
  .menu {
    width: 100%;
  }
}
</style>