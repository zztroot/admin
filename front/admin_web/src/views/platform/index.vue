<template>
    <template v-if="!isShow">
        <el-container>
            <el-header>
                <!-- 搜索框 -->
                <Header @isShow="openSetup" @searchData="startSearch"></Header>
            </el-header>
            <el-main>
                <!-- 主题内容 -->
                <Main
                    :tableData="tableData"
                    :total="total"
                    :loading="loading"
                    @isShow="modifyOpenSetup"
                    @isDelete="isDelete"
                    @handlerCurrentPage="handlerCurrentPage"
                    :currentPage="currentPage"
                ></Main>
            </el-main>
            <el-footer>
                <!-- 尾部 -->
            </el-footer>
        </el-container>
    </template>
    <template v-else>
        <router-view @isShow="closeSetup" :goodsOneData="goodsOneData"></router-view>
    </template>
</template>

<script setup lang='ts'>
import Main from './main.vue'
import Header from './header.vue'
import { getPlatformStatisticsList } from "@/api/platformStatistics"
import { computed, ref, watch } from "vue"
import { useRouter } from 'vue-router'
import { useStore } from "@/store/index"

let router = useRouter()
let store = useStore()
let tableData = ref()
let total = ref()
let loading = ref(true)
let goodsOneData = ref()
let isShow = computed(() => {
    return store.getters["getSetupShow"].platformSetup
})
let AuthorId = ref(0)
let PlatformId = ref(0)
let Name = ref("")
let currentPage = ref()

let modifyOpenSetup = (row: any) => {
    goodsOneData.value = row
    router.push("/platform/platformSetup")
}

let openSetup = (val: boolean) => {
    goodsOneData.value = undefined
    router.push("/platform/platformSetup")
}

let closeSetup = (val: boolean) => {
    goodsOneData.value = undefined
    router.push("/platform")
}

let getList = (page: string = "1", author_id: number = 0, platform_id: number = 0, name: string = "") => {
    getPlatformStatisticsList(page, "10", author_id, platform_id, name).then(r => {
        tableData.value = r.data.data.data
        total.value = r.data.data.total
        loading.value = false
    })
}

let startSearch = (name: string, authorId: number, platformId: number, startDate: string, endDate: string) => {
    getList("1", authorId, platformId, name)
    AuthorId.value = authorId
    PlatformId.value = platformId
    Name.value = name
    currentPage.value = 1
}

let isDelete = () => {
    getList(currentPage.value, AuthorId.value, PlatformId.value, Name.value)
}

let handlerCurrentPage = (page: number) => {
    currentPage.value = page
    getList(String(page), AuthorId.value, PlatformId.value, Name.value)
}

watch(isShow, (a, b) => {
    if (a == false) {
        // 查询数据
        getList(currentPage.value, AuthorId.value, PlatformId.value, Name.value)
    }
}, { "immediate": true })

</script>

<style scoped lang='scss'>
</style>