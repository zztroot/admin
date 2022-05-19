<template>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column label="ID" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span style="margin-left: 10px">{{ scope.row.id }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="Name" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span style="margin-left: 10px">{{ scope.row.id }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="Operations">
            <template #default="scope">
                <el-button size="small" @click="handleEdit(scope.$index)">Edit</el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.$index)">Delete</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>


<script setup lang='ts'>

const handleEdit = (index: number) => {
    console.log(index)
}
const handleDelete = (index: number) => {
    console.log(index)
}

const props = defineProps({
    tableData: {}
})
</script>


<style scoped lang='scss'>
</style>