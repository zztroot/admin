import { createRouter, createWebHistory } from 'vue-router';
import Layout from '@/layout/index.vue'

const routes: any = [
    {
        path: '/',
        name: 'login',
        component: () => import('@/views/login/login.vue'),
        meta: {
            title: '登录页面',
        }
    },
    {
        path: '/home',
        component: Layout,
        redirect: '/home',
        children: [
            {
                path: '/home',
                name: 'home',
                component: () => import('@/views/home/index.vue'),
                meta: {
                    title: '首页',
                }
            }
        ]
    },
    {
        path: '/goods',
        component: Layout,
        name: 'goods',
        children: [
            {
                path: '/goods',
                name: 'goods',
                component: () => import('@/views/goods/food.vue'),
                meta: {
                    title: '商品统计',
                },
                children: [
                    {
                        path: 'setup',
                        name: 'setup',
                        component: () => import('@/views/goods/foodSetup.vue'),
                        meta: {
                            title: '创建和修改',
                        }
                    }
                ]
            },
        ],
    },
    {
        path: '/platform',
        component: Layout,
        name: 'platform',
        children: [
            {
                path: '/platform',
                name: 'platform',
                component: () => import('@/views/platform/index.vue'),
                meta: {
                    title: '平台统计',
                },
                children: [
                    {
                        path: 'platformSetup',
                        name: 'platformSetup',
                        component: () => import('@/views/platform/setup.vue'),
                        meta: {
                            title: '创建和修改',
                        }
                    }
                ]
            }
        ],
    },
    {
        path: '/account',
        component: Layout,
        name: 'account',
        children: [
            {
                path: '/account',
                name: 'account',
                component: () => import('@/views/account/index.vue'),
                meta: {
                    title: '应付帐款',
                }
            }
        ],
    },
    {
        path: '/change',
        component: Layout,
        name: 'change',
        children: [
            {
                path: '/change',
                name: 'change',
                component: () => import('@/views/change/index.vue'),
                meta: {
                    title: '变动成本',
                }
            }
        ],
    },
    {
        path: '/fixed',
        component: Layout,
        name: 'fixed',
        children: [
            {
                path: '/fixed',
                name: 'fixed',
                component: () => import('@/views/fixed/index.vue'),
                meta: {
                    title: '固定成本',
                }
            }
        ],
    },
    {
        path: '/system',
        component: Layout,
        name: 'system',
        meta: {
            title: '系统管理',
        },
        children: [
            {
                path: '/user',
                name: 'user',
                component: () => import('@/views/system/user.vue'),
                meta: {
                    title: "用户管理"
                }
            },
            {
                path: '/role',
                name: 'role',
                component: () => import('@/views/system/role.vue'),
                meta: {
                    title: "角色管理"
                }
            }
        ],
    },
]

export default createRouter({
    history: createWebHistory(),
    routes
})
