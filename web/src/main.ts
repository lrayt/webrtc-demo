import './assets/main.css'

import {createApp} from 'vue'
import App from './App.vue'
import LayUI from '@layui/layui-vue'
import '@layui/layui-vue/lib/index.css'

import router from './router'

const app = createApp(App)
app.use(LayUI)
app.use(router)
app.mount('#app')
