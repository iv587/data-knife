import {createApp} from 'vue'
import Login from '@/pages/Login.vue'
import 'element-plus/dist/index.css'
import store from '@/store'
import '@/style.scss'
createApp(Login).use(store).mount("#app")