<template>
    <template v-if="!isShow">
        <el-container>
            <el-header>
                <!-- 搜索框 -->
                <FoodHeader @isShow="openSetup" @searchData="startSearch"></FoodHeader>
            </el-header>
            <el-main>
                <!-- 主题内容 -->
                <FoodMain
                    :tableData="tableData"
                    :total="total"
                    :loading="loading"
                    @isShow="modifyOpenSetup"
                    @isDelete="isDelete"
                    @handlerCurrentPage="handlerCurrentPage"
                    :currentPage="currentPage"
                ></FoodMain>
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
import FoodMain from './foodMain.vue'
import FoodHeader from './foodHeader.vue'
import { getGoodslist } from "@/api/goods"
import { computed, onMounted, reactive, ref, watch } from "vue"
import { useRouter } from 'vue-router'
import { useStore } from "@/store/index"
import { create } from '@/api/user'

let router = useRouter()
let store = useStore()
let tableData = ref()
let total = ref()
let loading = ref(true)
let goodsOneData = ref()
let isShow = computed(() => {
    return store.getters["getSetupShow"].goodsSetup
})
let AuthorId = ref(0)
let WarehouseId = ref(0)
let StartDate = ref("")
let EndDate = ref("")
let Name = ref("")
let currentPage = ref()

let modifyOpenSetup = (row: any) => {
    goodsOneData.value = row
    router.push("/goods/setup")
}

let openSetup = (val: boolean) => {
    goodsOneData.value = undefined
    router.push("/goods/setup")
}

let closeSetup = (val: boolean) => {
    goodsOneData.value = undefined
    router.push("/goods")
}

let getFoodList = (page: string = "1", author_id: number = 0, warehouse_id: number = 0, start_date: string = "", end_date: string = "", name: string = "") => {
    getGoodslist(page, "10", author_id, warehouse_id, start_date, end_date, name).then(r => {
        total.value = r.data.data.total
        setTimeout(() => {
            tableData.value = r.data.data.data
            loading.value = false
        }, 100);
    })
}

let startSearch = (name: string, authorId: number, warehouseId: number, startDate: string, endDate: string) => {
    getFoodList("1", authorId, warehouseId, startDate, endDate, name)
    AuthorId.value = authorId
    WarehouseId.value = warehouseId
    StartDate.value = startDate
    EndDate.value = endDate
    Name.value = name
    currentPage.value = 1
}

let isDelete = () => {
    getFoodList(currentPage.value, AuthorId.value, WarehouseId.value, StartDate.value, EndDate.value, Name.value)
}

let handlerCurrentPage = (page: number) => {
    currentPage.value = page
    getFoodList(String(page), AuthorId.value, WarehouseId.value, StartDate.value, EndDate.value, Name.value)
}

watch(isShow, (a, b) => {
    if (a == false) {
        // 查询数据
        getFoodList(currentPage.value, AuthorId.value, WarehouseId.value, StartDate.value, EndDate.value, Name.value)
    }
}, { "immediate": true })

</script>

<style scoped lang='scss'>
</style>