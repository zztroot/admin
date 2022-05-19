// store.ts
import { InjectionKey } from 'vue'
import { createStore, useStore as baseUseStore, Store } from 'vuex'

export interface State {
    count: number,
    collapse: boolean,
    roleDrawer: boolean,
    token: string,
    isSelect: number,
    auth: any,
    isCommit: number,
    isQueryUserList: number,
    userHandlerSearch: UserSearch,
    setupShow: Setup
}

export interface UserSearch {
    userName: string,
    sexId: number | undefined,
    isSearch: number
}

// 创建和修改页面
export interface Setup {
    goodsSetup?: boolean,
    platformSetup?: boolean
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state: {
        count: 0,
        collapse: false,
        roleDrawer: false,
        token: "",
        isSelect: 0,
        auth: [],
        isCommit: 0,
        isQueryUserList: 0,
        userHandlerSearch: {
            userName: "",
            sexId: 0,
            isSearch: 0
        },
        setupShow: {
            goodsSetup: false,
            platformSetup: false
        }
    },
    mutations: {
        setCount(state: State, count: number) {
            state.count = count;
        },
        // 设置collapse
        setCollapse: (state: State, collapse: boolean) => {
            state.collapse = collapse
        },
        // 设置roleDrawer
        setRoleDrawer: (state: State, roleDrawer: boolean) => {
            state.roleDrawer = roleDrawer
        },
        // 设置token
        setToken: (state: State, token: string) => {
            state.token = token
        },
        // 设置提交为true
        setIsSelect: (state: State, isSelect: number) => {
            state.isSelect = isSelect
        },
        // 设置auth
        setAuth: (state: State, auth: any) => {
            state.auth = auth
        },
        // 设置提交为true
        setIsCommit: (state: State, isCommit: number) => {
            state.isCommit = isCommit
        },
        // 设置再次查询用户
        setIsQueryUserList: (state: State, isQueryUserList: number) => {
            state.isQueryUserList = isQueryUserList
        },
        // 设置用户查询数据
        setUserHandlerSearch: (state: State, userHandlerSearch: UserSearch) => {
            let temp = state.userHandlerSearch
            temp.userName = userHandlerSearch.userName
            temp.sexId = userHandlerSearch.sexId
            temp.isSearch = userHandlerSearch.isSearch
            state.userHandlerSearch = temp
        },
        // 设置创建和修改页面是否展示
        setSetupShow: (state: State, setupShow: Setup) => {
            let temp = state.setupShow
            temp.goodsSetup = setupShow.goodsSetup
            temp.platformSetup = setupShow.platformSetup
            state.setupShow = temp
        }
    },
    getters: {
        getCount(state: State) {
            return state.count
        },
        // 获取collapse
        getCollapse: (state: State) => {
            return state.collapse
        },
        // 获取roleDrawer
        getRoleDrawer: (state: State) => {
            return state.roleDrawer
        },
        // 获取token
        getToken: (state: State) => {
            return state.token
        },
        // 获取isSelect
        getIsSelect: (state: State) => {
            return state.isSelect
        },
        // 获取auth
        getAuth: (state: State) => {
            return state.auth
        },
        getIsCommit: (state: State) => {
            return state.isCommit
        },
        getIsQueryUserList: (state: State) => {
            return state.isQueryUserList
        },
        getUserHandlerSearch: (state: State): UserSearch => {
            return state.userHandlerSearch
        },
        getSetupShow: (state: State): Setup => {
            return state.setupShow
        }
    }
})

// 定义自己的 `useStore` 组合式函数
export function useStore() {
    return baseUseStore(key)
}