<template>
    <el-table
        :data="tableData"
        style="width: 100%"
        v-loading="loading"
        :cell-style="cellStyle"
        :header-cell-style="headerCellStyle"
    >
        <el-table-column fixed prop label="ID" width="40">
            <template #default="scope">
                <span>{{ scope.$index + 1 }}</span>
            </template>
        </el-table-column>
        <el-table-column fixed prop="goods.goods_name" label="名称" width="250" />
        <el-table-column prop="goods.goods_code" label="编号" width="160" />
        <el-table-column prop="goods.goods_sale_date" label="销售日期" width="120" />
        <el-table-column prop="warehouse" label="仓库" width="120" />
        <el-table-column prop="author" label="创建者" width="120" />
        <el-table-column prop="goods.goods_unit" label="单位" width="120" />
        <el-table-column prop="goods.goods_cost_price" label="成本价" width="120" />
        <el-table-column prop="goods.goods_supply_price" label="供货价" width="120" />
        <el-table-column prop="goods.goods_sale_number" label="销售数量" width="120" />
        <el-table-column prop="goods.goods_return_number" label="客退数量" width="120" />
        <el-table-column prop="goods.goods_link_return_number" label="链路客退数量" width="120" />
        <el-table-column prop="income.sales_volume" label="销售额" width="100" />
        <el-table-column prop="income.return_volume" label="客退额" width="100" />
        <el-table-column prop="income.link_return_volume" label="链路客退额" width="100" />
        <el-table-column prop="income.sales_volume_after_return" label="扣退后销售额" width="120" />
        <el-table-column prop="cost.goods_cost" label="商品成本" width="100" />
        <el-table-column prop="cost.return_cost" label="客退成本" width="100" />
        <el-table-column prop="cost.link_return_cost" label="链路客退成本" width="120" />
        <el-table-column prop="cost.goods_cost_after_return" label="扣退后产品成本" width="125" />
        <el-table-column prop="profit.gross_profit" label="毛利额" width="100" />
        <el-table-column prop="profit.gross_profit_rate" label="毛利率" width="100" />
        <el-table-column prop="return_rate" label="客退率" width="100" />
        <el-table-column prop="goods.create_time" label="创建日期" width="180" />
        <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
                <el-button size="small" type="primary" @click="handleEdit(scope.row)">修改</el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.row.goods.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>


<script setup lang='ts'>
import { deleteGoods } from "@/api/goods"
import { ElMessage, ElMessageBox } from "element-plus"

let emits = defineEmits(["isShow", "isDelete"])

const handleEdit = (row: any) => {
    emits("isShow", row)
}

let cellStyle = function (param: any = {}) {
    const { column } = param
    if (column.label == '销售数量' || column.label == '毛利额') {
        return { color: "rgb(10, 94, 28)" }
    }
    if (column.label == '扣退后销售额' || column.label == "扣退后产品成本") {
        return { "color": "rgb(233, 89, 149)" }
    }
    return ""
}

let headerCellStyle = (param: any): any => {
    const { column } = param
    if (column.label == '销售数量' || column.label == '毛利额') {
        return { color: "rgb(10, 94, 28)" }
    }
    if (column.label == '扣退后销售额' || column.label == "扣退后产品成本") {
        return { "color": "rgb(233, 89, 149)" }
    }
    return ""
}

const handleDelete = (index: number) => {
    ElMessageBox.confirm('确定删除该商品吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        deleteGoods(index).then(r => {
            if (r.data.code == 0) {
                ElMessage.success("删除成功")
                emits("isDelete", true)
            }
        })
    }).catch(() => {

    })
}

defineProps(["tableData", "loading"])
</script>


<style lang='scss' scoped>
</style>