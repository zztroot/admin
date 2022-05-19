<template>
    <el-pagination
        background
        layout="total, prev, pager, next"
        :total="total"
        :hide-on-single-page="true"
        :current-page="currentPage"
        @current-change="currentChange"
    />
</template>


<script setup lang='ts'>
import { ref } from "vue"

defineProps(["total", "currentPage"])
let emits = defineEmits(["getPage"])

let currentChange = (page: number) => {
    emits("getPage", page)
}

</script>

<style scoped lang='scss'>
.el-pagination {
    margin-top: 20px;
    float: right;
}
</style>