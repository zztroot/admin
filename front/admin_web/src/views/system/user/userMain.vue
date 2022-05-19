<template>
    <DefaultTable :tableData="tableData" :loading="loading" :roleList="roleList"></DefaultTable>
    <template v-if="total > 10">
        <Pagination :total="total" @getPage="getPage"></Pagination>
    </template>
</template>

<script setup lang='ts'>
import DefaultTable from '@/views/system/user/userTable.vue'
import Pagination from '@/components/pagination/default.vue'
import { computed, onMounted, reactive, ref, watch } from "vue"
import { getlist } from "@/api/user"
import { useStore } from "@/store/index"
import type { UserSearch } from "@/store/index"

let store = useStore()
let tableData = ref()
let total = ref()
let isQueryUser = computed(() => {
    return store.getters["getIsQueryUserList"]
})
let loading = ref(true)
let isSearchUser = computed(() => {
    return store.getters["getUserHandlerSearch"].isSearch
})

let props = defineProps<{
    roleList?: any,
}>()

onMounted(() => {
    getUserList("1", "", 0)
})

let getUserList = (page: string, user_name: string, sex: number) => {
    // 获取用户列表
    getlist(page, "10", user_name, sex).then(r => {
        tableData.value = r.data.data.data
        total.value = r.data.data.total
        loading.value = false
    })
}

let getPage = (page: number) => {
    getUserList(String(page), "", 0)
}

// 监听isQueryUser
watch(isQueryUser, (a, b) => {
    if (a != 0) {
        setTimeout(() => {
            getUserList("1", "", 0)
        }, 100);
        store.commit("setIsQueryUserList", 0)
    }
}, { "immediate": true })

// 监听用户搜索
watch(isSearchUser, (a, b) => {
    if (a != 0) {
        let search = store.getters["getUserHandlerSearch"]
        getUserList("1", search.userName, search.sexId)
        let searchUser: UserSearch = {
            userName: "",
            sexId: 0,
            isSearch: 0
        }
        store.commit("setUserHandlerSearch", searchUser)
    }
}, { "immediate": true })
</script>


<style scoped lang='scss'>
</style>