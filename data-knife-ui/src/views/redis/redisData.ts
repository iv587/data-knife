import {listConnection} from "@/api/connection";
import {redisDbsApi, redisKeysApi} from "@/api/redis";
import {onMounted, ref} from "vue";
import {RedisKeyListRes} from "@/type";
import RedisDataUpdatePanel from "@/components/redis/RedisDataUpdatePanel.vue";

const dataTagTypeMap = new Map<string, string>()
dataTagTypeMap.set('list', 'waring')
dataTagTypeMap.set('set', 'info')
dataTagTypeMap.set('hash', '')
dataTagTypeMap.set('zset', 'danger')
dataTagTypeMap.set('string', 'success')

export type RedisDataUpdatePanelType =  InstanceType<typeof RedisDataUpdatePanel> | null


export function useRedisData(id: number) {


    const pattern = ref<{
        left: number
        right: number
    }>({
        left: 1,
        right: 1
    })

    const keyPattern = ref('')


    const dbList = ref<number[]>()

    const selectDbNum = ref<number>(0)

    const dataFetchLoading = ref(false)

    const keysRes = ref<RedisKeyListRes>({
        list: [],
        total: 0,
        max: 0
    })

    const redisDbHandler = async () => {
        const res = await listConnection()
        if (res.data.list.length > 0) {
            const dbsRes = await redisDbsApi({id,})
            dbList.value = dbsRes.data.list
            selectDbNum.value = dbList.value[0]
            await listKeysHandler()
        }
    }

    const onSelect = async (dbNum: number) => {
        selectDbNum.value = dbNum
        await listKeysHandler()
    }

    const listKeysHandler = async () => {
        dataFetchLoading.value = true
        try {
            let searchKey = keyPattern.value
            const {left, right} = pattern.value
            if (left == 1) {
                searchKey = '*' + searchKey
            }
            if (right == 1) {
                searchKey =  searchKey + '*'
            }
            const res = await redisKeysApi({id, dbNo: selectDbNum.value, keyPattern: searchKey})
            const {list, max, total} = res.data
            keysRes.value.list = list
            keysRes.value.max = max
            keysRes.value.total = total
        } finally {
            dataFetchLoading.value = false
        }
    }




    const showDashboardHandler = (emit: any, show: boolean) => {
        emit('close')
    }

    const openUpdatePanelHandler = (redisDataUpdatePanel: any, key?: string) => {
        redisDataUpdatePanel.openHandler(key)
    }

    onMounted(async () => {
        await redisDbHandler()
    })

    return {
        pattern, keyPattern, dbList, selectDbNum, dataFetchLoading, keysRes,
        redisDbHandler, onSelect, listKeysHandler, showDashboardHandler,
        openUpdatePanelHandler, dataTagTypeMap
    }
}