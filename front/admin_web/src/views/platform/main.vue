<template>
    <Table :tableData="tableData" :loading="loading" @isShow="isShow" @isDelete="isDelete"></Table>
    <template v-if="total > 10">
        <Pagination :total="total" @getPage="getPage" :currentPage="currentPage"></Pagination>
    </template>
</template>


<script setup lang='ts'>
import Table from './table.vue'
import Pagination from '@/components/pagination/default.vue'
import { reactive, ref } from 'vue'

let props = defineProps(["tableData", "total", "loading", "currentPage"])

let emtis = defineEmits(["isShow", "isDelete", "handlerCurrentPage"])

let getPage = (page: number) => {
    emtis("handlerCurrentPage", page)
}

let isShow = (row: any) => {
    emtis("isShow", row)
}

let isDelete = () => {
    emtis("isDelete", true)
}
</script>


<style scoped lang='scss'>
</style>