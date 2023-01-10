import {createPinia } from "pinia"
import {createPersistedState } from 'pinia-persistedstate-plugin'
import type {App} from 'vue'

const pinia = createPinia()
pinia.use(createPersistedState())

export default {
    install(app: App) {
        app.use(pinia)
    }
}