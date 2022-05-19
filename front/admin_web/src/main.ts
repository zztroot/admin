import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router/index'
import { store, key } from '@/store/index'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as Icons from '@element-plus/icons-vue'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

const app = createApp(App)
app.use(router).use(store, key).use(ElementPlus, { locale: zhCn })
app.mount('#app')

// 注册全局图标组件
Object.keys(Icons).forEach((key) => {
    app.component(key, Icons[key as keyof typeof Icons])
})

router.beforeEach((to, from, next) => {
    // to 是从哪里来
    // from 去哪里
    // next  next() 代表放行  next('/login') 强制跳回login页面
    // 获取token 
    let user = localStorage.getItem('user')
    // 两种没有登录的情况 user为null
    if (user == null && to.path === '/') return next()
    if (user == null && to.path !== '/') return next('/')
    // 两种登录的情况 user不为null
    let users = JSON.parse(String(user))
    // 1.当前页面是/ 登录页面，并且token不为空，直接进入首页
    if (to.path == '/' && users.token !== '') return next('/home')
    // 2.不论那个页面，只要token存在，直接放行
    if (users.token !== '') return next()
})
