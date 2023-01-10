import {App} from "vue";
import RedisIcon from '@/icon/RedisIcon.vue'
import ConnectionIcon from '@/icon/ConnectionIcon.vue'
import RedisMenuInfoIcon from '@/icon/RedisMenuInfoIcon.vue'
import RedisDataMenuIcon from '@/icon/RedisDataMenuIcon.vue'

export default {
    install: (app: App) => {
        app.component('redis-icon', RedisIcon)
        app.component('connection-icon', ConnectionIcon)
        app.component('redis-menu-info-icon', RedisMenuInfoIcon)
        app.component('redis-data-menu-icon', RedisDataMenuIcon)
    },
}