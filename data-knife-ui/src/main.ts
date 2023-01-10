import { createApp } from 'vue'
import '@/style.scss'
import App from '@/App.vue'
import store from '@/store'
import Router from '@/router'
import MyIcon from '@/icon'
import 'element-plus/dist/index.css'



createApp(App)
    .use(Router)
    .use(store)
    .use(MyIcon)
    .mount('#app')
