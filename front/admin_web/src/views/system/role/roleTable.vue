<template>
    <DefaultButton name="创建" @click="dialogVisible = true" types="primary"></DefaultButton>
    <BaseDialog
        :dialogVisible="dialogVisible"
        @dialogClost="dialogClost"
        @dialogCreate="dialogCreate"
        title="创建角色"
        buttonName="创建"
        width="20%"
    >
        <template #dialogMain>
            <el-input v-model="input" placeholder="请输入角色名称" />
        </template>
    </BaseDialog>
    <el-table :data="tableData" style="width: 100%" @row-click="selectRole">
        <el-table-column label="ID" width="120">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.$index + 1 }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="角色名称" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.role_name }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="创建时间" width="200">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.create_time }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="操作">
            <template #default="scope">
                <!-- <el-button
                    size="small"
                    type="primary"
                    @click="handleEdit(scope.row.id, scope.row.role_name)"
                >修改</el-button>-->
                <el-button size="small" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script setup lang='ts'>
import { computed, onMounted, reactive, ref } from 'vue'
import { useStore } from '@/store/index'
import Drawer from '@/components/drawer/rightDrawer.vue'
import DefaultButton from '@/components/button/defaultButton.vue'
import BaseDialog from "@/components/dialog/baseDialog.vue"
import { createRole, deleteRole } from "@/api/role"
import type { httpRequests } from "@/api/api"
import { ElMessage, ElMessageBox } from 'element-plus'

let props = defineProps<{
    tableData?: any,
}>()
const dialogVisible = ref(false)
let input = ref("")

let emits = defineEmits(["roleData", "roleCreate"])

// 编辑角色
const handleEdit = (index: number, name: string) => {

}

const dialogClost = () => {
    dialogVisible.value = false
}

// 删除角色
const handleDelete = (index: number) => {
    ElMessageBox.confirm('确定删除吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        deleteRole(index).then(r => {
            if (r.data.code == 0) {
                ElMessage.success("删除成功")
                emits("roleCreate", true)
            }
        })
    }).catch(() => {

    })
}

// 角色选择
const selectRole = (row: any) => {
    emits("roleData", row)
}

// 创建角色
const dialogCreate = () => {
    createRole(input.value).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("创建成功")
            emits("roleCreate", true)
            input.value = ""
        }
    })
    dialogVisible.value = false
}
</script>


<style scoped lang='scss'>
</style>