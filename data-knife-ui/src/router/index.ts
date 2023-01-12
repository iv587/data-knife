import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import auth from "@/utils/auth";
import NProgress from "nprogress";
import 'nprogress/nprogress.css'

const routes = (device: string):RouteRecordRaw[] => {
    return [
        {
            path: '/',
            name: 'MainView',
            component: () => import(`@/views/${device}Main.vue`),
            redirect: '/connections',
            children: [
                {
                    path: 'connections',
                    name: 'ConnectionView',
                    component: () => import('@/views/Connection.vue')
                },
                {
                    path: 'redis/:id',
                    name: 'RedisView',
                    component: () => import('@/views/redis/RedisView.vue')
                },
                {
                    path: 'redis/info/:id',
                    name: 'RedisInfoView',
                    component: () => import('@/views/redis/RedisInfoView.vue')
                },
            ]
        },
    ]
}



const genRouter = (device: string) => {
    const router = createRouter({
        history: createWebHashHistory(),
        routes: routes(device),
    })

    router.beforeEach((to, from, next) => {
        NProgress.start()
        const token = auth.getToken()
        if (!token) {
            auth.toLogin()
        } else {
            next()
        }
    })
    router.afterEach(() => {
        NProgress.done()
    })
    return router
}

export default genRouter

