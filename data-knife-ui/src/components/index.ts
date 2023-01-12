import {App} from "vue";
import WebContainer from '@/components/WebContainer.vue'

export default {
    install: (app: App) => {
        app.component('web-container', WebContainer)
    },
}