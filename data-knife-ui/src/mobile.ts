import { createApp } from 'vue'
import App from '@/App.vue'
import store from '@/store'
import MyIcon from '@/icon'
import 'element-plus/dist/index.css'
import genRouter from "@/router";
import '@/style.scss'

createApp(App)
    .use(store)
    .use(genRouter('M'))
    .use(MyIcon)
    .mount('#app')
