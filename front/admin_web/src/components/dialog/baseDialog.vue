<template>
    <el-dialog v-model="dialogVisible" :title="title" :width="width" :before-close="handleClose">
        <span>
            <slot name="dialogMain"></slot>
        </span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogClost()">关闭</el-button>
                <el-button type="primary" @click="dialogCreate()">{{ buttonName }}</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang='ts'>
import { ElMessageBox } from "element-plus"
import { ref } from "vue"


defineProps(["dialogVisible", "title", "buttonName", "width"])

let emits = defineEmits(["dialogClost", "dialogCreate"])

const dialogClost = () => {
    emits("dialogClost", true)
}

const dialogCreate = () => {
    emits("dialogCreate", true)
}

const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定退出窗口吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        done()
        emits("dialogClost", true)
    }).catch(() => {

    })
}
</script>


<style scoped lang='scss'>
</style>