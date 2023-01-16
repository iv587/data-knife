<template>
  <el-menu :class="menuCls?? 'menu'"
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

const collapse= computed(() => {
  return menuStore.collapse
})

const onClick = (index: string) => {
  menuStore.setActiveIndex(index)
  emit('click')
}

</script>

<style lang="scss">
.menu:not(.el-menu--collapse) {
  width: 250px;
}
.menu1:not(.el-menu--collapse) {
  width: 100%;
}
</style>