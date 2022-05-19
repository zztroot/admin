<template>
    <el-row :gutter="25">
        <el-col :span="5">
            <div class="grid-content bg-purple">
                <el-input v-model="name" placeholder="商品名称">
                    <template #prepend>名称</template>
                </el-input>
            </div>
        </el-col>
        <el-col :span="5">
            <el-select v-model="authorId" placeholder="选择创建者">
                <el-option v-for="val in authorList" :label="val.name" :value="val.id" />
            </el-select>
        </el-col>
        <el-col :span="5">
            <el-select v-model="warehouseId" placeholder="选择仓库">
                <el-option
                    v-for="val in warehouseList"
                    :label="val.warehouse_name"
                    :value="val.id"
                />
            </el-select>
        </el-col>
        <el-col :span="6">
            <el-date-picker
                v-model="valueDate"
                type="daterange"
                range-separator="To"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
            />
        </el-col>
        <DefaultButton name="重置" @click="reset"></DefaultButton>
        <DefaultButton name="搜索" @click="search" types="primary"></DefaultButton>
        <div class="createUser">
            <DefaultButton name="创建" @click="create" types="primary"></DefaultButton>
        </div>
    </el-row>
</template>


<script setup lang='ts'>
import DefaultInput from '@/components/input/defaultInput.vue'
import DefaultButton from '@/components/button/defaultButton.vue'
import { onMounted, ref } from "vue"
import { dateFormat } from "@/util/time"
import { getWarehouseList } from "@/api/warehouse"
import { getUserAll } from "@/api/user"

let name = ref("")
let authorId = ref()
let warehouseId = ref()
let valueDate = ref("")
let authorList = ref()
let warehouseList = ref()

let emits = defineEmits(["isShow", "searchData"])

onMounted(() => {
    getWarehouseList().then(r => {
        warehouseList.value = r.data.data
    })
    getUserAll().then(r => {
        authorList.value = r.data.data
    })
})

let search = () => {
    if (authorId.value == "") {
        authorId.value = 0
    }
    if (warehouseId.value == "") {
        warehouseId.value = 0
    }
    let start: string = ""
    let end: string = ""
    if (valueDate.value != "") {
        start = dateFormat(valueDate.value[0])
        end = dateFormat(valueDate.value[1])
    }
    emits("searchData", name.value, authorId.value, warehouseId.value, start, end)
    if (authorId.value == 0) {
        authorId.value = ""
    }
    if (warehouseId.value == 0) {
        warehouseId.value = ""
    }
}

let reset = () => {
    name.value = ""
    authorId.value = ""
    warehouseId.value = ""
    valueDate.value = ""
    emits("searchData", "", 0, 0, "", "")
}

let create = () => {
    emits("isShow", false)
}
</script>

<style scoped lang='scss'>
.createUser {
    position: absolute;
    right: 0px;
}
.el-col-5 {
    flex: none;
}
</style>