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
            <el-select v-model="warehouseId" placeholder="选择平台">
                <el-option v-for="val in platfrom" :label="val.platform_name" :value="val.id" />
            </el-select>
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
import { getPlatfromList } from "@/api/platfrom"
import { getUserAll } from "@/api/user"

let name = ref("")
let authorId = ref()
let warehouseId = ref()
let authorList = ref()
let platfrom = ref()

let emits = defineEmits(["isShow", "searchData"])

onMounted(() => {
    getPlatfromList().then(r => {
        platfrom.value = r.data.data
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