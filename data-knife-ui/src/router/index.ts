import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import auth from "@/utils/auth";
import NProgress from "nprogress";
import 'nprogress/nprogress.css'

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'MainView',
        component: () => import('@/views/MainView.vue'),
        redirect: '/connections',
        children: [
            {
                path: 'connections',
                name: 'ConnectionView',
                component: () => import('@/views/ConnectionView.vue')
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
    {
        path: '/login',
        name: 'LoginView',
        component: () => import('@/views/LoginView.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    NProgress.start()
    const token = auth.getToken()
    if (!token) {
        if (to.name == 'LoginView') {
            next()
        } else {
            next({name: 'LoginView'})
        }
    } else {
        if (to.name == 'LoginView') {
            next({path: '/'})
        } else {
            next()
        }
    }
})

router.afterEach(() => {
    NProgress.done()
})

export default router

