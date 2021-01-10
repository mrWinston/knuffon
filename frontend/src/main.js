import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'

const app = createApp(App).use(router).use(store)

if (process.env.NODE_ENV === 'development') {
  console.log("DEV ENV!!!")
  if ('__VUE_DEVTOOLS_GLOBAL_HOOK__' in window) {
    window.__VUE_DEVTOOLS_GLOBAL_HOOK__.Vue = app
  }
}

app.mount('#app')
