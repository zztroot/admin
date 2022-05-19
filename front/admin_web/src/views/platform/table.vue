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
        <el-table-column fixed prop="goods.goods_model.goods_name" label="名称" width="250" />
        <el-table-column prop="goods.goods_model.goods_code" label="编号" width="160" />
        <el-table-column prop="goods.goods_model.goods_sale_date" label="销售日期" width="120" />
        <el-table-column prop="goods.warehouse" label="仓库" width="120" />
        <el-table-column prop="plat_form.platform_str" label="平台" width="120" />
        <el-table-column prop="plat_form.author" label="创建者" width="120" />
        <el-table-column prop="goods.goods_model.goods_unit" label="单位" width="100" />
        <el-table-column prop="goods.goods_model.goods_cost_price" label="成本价" width="120" />
        <el-table-column prop="goods.goods_model.goods_supply_price" label="供货价" width="120" />
        <el-table-column prop="goods.goods_model.goods_sale_number" label="销售数量" width="120" />
        <el-table-column prop="goods.goods_cost_total" label="商品成本" width="120" />
        <el-table-column prop="company.company_sale_volume" label="公司销售额" width="120" />
        <el-table-column prop="company.company_one_gross_profit" label="公司单袋毛利额" width="130" />
        <el-table-column prop="company.company_gross_profit" label="公司毛利额" width="100" />
        <el-table-column prop="company.company_profit_rate" label="公司毛利率" width="100" />
        <el-table-column prop="plat_form.platform_model.platform_price" label="平台售价" width="120" />
        <el-table-column prop="plat_form.platform_sale_volume" label="平台销售额" width="100" />
        <el-table-column prop="plat_form.platform_gross_profit" label="平台毛利额" width="100" />
        <el-table-column prop="plat_form.platform_profit_rate" label="平台毛利率" width="120" />
        <el-table-column prop="plat_form.platform_model.platform_comments" label="备注" width="180" />
        <el-table-column prop="plat_form.platform_model.create_time" label="创建日期" width="180" />
        <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
                <el-button size="small" type="primary" @click="handleEdit(scope.row)">修改</el-button>
                <el-button
                    size="small"
                    type="danger"
                    @click="handleDelete(scope.row.plat_form.platform_model.id)"
                >删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>


<script setup lang='ts'>
import { deletePlatformStatistics } from "@/api/platformStatistics"
import { ElMessage, ElMessageBox } from "element-plus"

let emits = defineEmits(["isShow", "isDelete"])

const handleEdit = (row: any) => {
    emits("isShow", row)
}

let cellStyle = function (param: any = {}) {
    const { column } = param
    if (column.label == '销售数量' || column.label == '平台') {
        return { color: "rgb(10, 94, 28)" }
    }
    if (column.label == '平台销售额' || column.label == "公司销售额") {
        return { "color": "rgb(233, 89, 149)" }
    }
    return ""
}

let headerCellStyle = (param: any): any => {
    const { column } = param
    if (column.label == '销售数量' || column.label == '平台') {
        return { color: "rgb(10, 94, 28)" }
    }
    if (column.label == '平台销售额' || column.label == "公司销售额") {
        return { "color": "rgb(233, 89, 149)" }
    }
    return ""
}

const handleDelete = (index: number) => {
    ElMessageBox.confirm('确定删除该商品吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        deletePlatformStatistics(index).then(r => {
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