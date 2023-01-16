import {computed, watch} from "vue";
import useMenuStore from "@/store/menu";
import {useRoute, useRouter} from "vue-router"

export function useMenuData() {
    const route = useRoute()
    const router = useRouter()
    const currentRoute = router.currentRoute
    const menuStore = useMenuStore()

    const list = computed(() => {
        return menuStore.menuList
    })

    watch(currentRoute, (newRoute, oldRoute) => {
        menuStore.setActiveIndex(newRoute.path)
    })

    return {list, route}
}

