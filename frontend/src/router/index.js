import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    redirect: '/redis/connections',
    component: () => import('@/views/Main'),
    children: [
      {
        path: '/redis/connections',
        name: 'RedisConnection',
        component: () => import('@/views/redis/RedisServer')
      },
      {
        path: '/redis/db',
        name: 'RedisData',
        component: () => import('@/views/redis/RedisData')
      },
      {
        path: '/redis/info',
        name: 'RedisInfo',
        component: () => import('@/views/redis/RedisInfo')
      },
      {
        path: '/user/setting',
        name: 'UserSetting',
        component: () => import('@/views/redis/UserSetting')
      }
    ]
  },
  {
    name: 'Login',
    path: '/login',
    component: () => import('@/views/Login')
  }
]

const router = new VueRouter({
  scrollBehavior: () => ({ y: 0 }),
  routes
})

export default router
