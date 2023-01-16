import useMenuStore from "@/store/menu";
import {useRouter} from "vue-router";
import {onMounted, Ref, ref} from "vue";
import {Connection} from "@/type";
import {listConnection} from "@/api/connection";
import ConnectionUpdatePanel from "@/components/connection/ConnectionUpdatePanel.vue";

export function useConnection() {
    const menuStore = useMenuStore()

    const router = useRouter()


    const list = ref<Connection[]>([])

    const loading = ref(false)


    const listConnectionHandler = async () => {
        loading.value = true
        try {
            const res = await listConnection()
            const data = res.data
            list.value = data.list
        } finally {
            loading.value = false
        }
    }

    const infoHandler = async (conn: Connection) => {
        menuStore.updateMenuItem({
            id: conn.id + '',
            index: `/redis/${conn.id}`,
            title: conn.name,
            icon: 'redis-icon',
        })
        await router.push({path: `/redis/${conn.id}`})
    }

    const onUpdateSuccessHandler = async (id?: number) => {
        await listConnectionHandler()
        if (id != null && id > 0) {
            for (const i in list.value) {
                if (list.value[i].id == id) {
                    const conn = list.value[i]
                    menuStore.updateMenuItem({
                        id: conn.id + '',
                        index: `/redis/${conn.id}`,
                        title: conn.name,
                        icon: 'redis-icon',
                    }, true)
                }
            }
        }

    }

    const openConnectionUpdatePanelHandler = (connectionUpdatePanelRef: InstanceType<typeof ConnectionUpdatePanel>, id?: number) => {
        connectionUpdatePanelRef.openHandler()
    }

    onMounted(async () => {
        await listConnectionHandler()
    })

    return {
        list,
        loading,
        infoHandler,
        onUpdateSuccessHandler,
        openConnectionUpdatePanelHandler,
    }
}