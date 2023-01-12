import { createApp } from 'vue'
import '@/style.scss'
import App from '@/App.vue'
import store from '@/store'
import genRouter from '@/router'
import MyIcon from '@/icon'
import components from '@/components'
import 'element-plus/dist/index.css'



createApp(App)
    .use(components)
    .use(genRouter(''))
    .use(store)
    .use(MyIcon)
    .mount('#app')