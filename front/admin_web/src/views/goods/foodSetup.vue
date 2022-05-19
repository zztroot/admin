<template>
    <el-container>
        <el-main>
            <el-form
                ref="ruleFormRef"
                :model="ruleForm"
                :rules="rules"
                label-width="120px"
                class="demo-ruleForm"
                :size="formSize"
            >
                <el-container>
                    <el-aside width="50%">
                        <el-form-item label="商品名称" prop="name">
                            <el-input v-model="ruleForm.name" />
                        </el-form-item>
                        <el-form-item label="销售日期" prop="date">
                            <el-date-picker
                                v-model="ruleForm.date"
                                type="date"
                                placeholder="请选择销售日期"
                                style="width: 50%"
                            />
                        </el-form-item>
                        <el-form-item label="商品单位" prop="unit">
                            <el-input v-model="ruleForm.unit" />
                        </el-form-item>
                        <el-form-item label="成本价" prop="cost_price">
                            <el-input v-model="ruleForm.cost_price" />
                        </el-form-item>
                        <el-form-item label="选择仓库" prop="warehouse">
                            <el-select v-model="ruleForm.warehouse" placeholder="选择仓库">
                                <el-option
                                    v-for="val in warehouse"
                                    :label="val.warehouse_name"
                                    :value="val.id"
                                />
                            </el-select>
                            <span class="z_a">
                                没有仓库？
                                <el-button type="text" @click="dialogVisible = true">去添加</el-button>
                            </span>
                        </el-form-item>
                    </el-aside>
                    <el-main>
                        <el-form-item label="供货价格" prop="supply_price">
                            <el-input v-model="ruleForm.supply_price" />
                        </el-form-item>
                        <el-form-item label="销售数量" prop="sale_number">
                            <el-input v-model="ruleForm.sale_number" />
                        </el-form-item>
                        <el-form-item label="退货数量" prop="return_number">
                            <el-input v-model="ruleForm.return_number" />
                        </el-form-item>
                        <el-form-item label="链路客退数量" prop="link_return_number">
                            <el-input v-model="ruleForm.link_return_number" />
                        </el-form-item>
                    </el-main>
                </el-container>
                <el-form-item class="zt_commit">
                    <el-button type="info" @click="returnPage">上一页</el-button>
                    <el-button @click="resetForm(ruleFormRef)">重置</el-button>
                    <el-button type="primary" @click="submitForm(ruleFormRef)">提交</el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
    <BaseDialog
        :dialogVisible="dialogVisible"
        @dialogClost="dialogClost"
        @dialogCreate="dialogCreate"
        title="添加仓库"
        buttonName="添加"
        width="20%"
    >
        <template #dialogMain>
            <el-input v-model="warehouseInput" placeholder="请输入仓库名称" />
        </template>
    </BaseDialog>
</template>


<script setup lang='ts'>
import DefaultButton from '@/components/button/defaultButton.vue'
import { ElMessageBox, ElMessage } from "element-plus"
import { onMounted, reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import BaseDialog from "@/components/dialog/baseDialog.vue"
import type { SaveGoodsParams } from "@/api/goods"
import { dateFormat } from "@/util/time"
import { saveGoods, modifyGoods } from "@/api/goods"
import { getWarehouseList, saveWarehouse } from "@/api/warehouse"

const dialogVisible = ref(false)
const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const warehouseInput = ref("")
let ruleForm = reactive({
    name: '',
    date: '',
    unit: '',
    cost_price: 0,
    warehouse: '',
    supply_price: 0,
    sale_number: 0,
    return_number: 0,
    link_return_number: 0
})
let warehouse = ref()
let goodsId = ref(0)

let emits = defineEmits(["isShow"])

let props = defineProps(["goodsOneData"])

onMounted(() => {
    getWarehouses()
    if (props.goodsOneData != undefined) {
        ruleForm.name = props.goodsOneData.goods.goods_name
        ruleForm.date = props.goodsOneData.goods.goods_sale_date
        ruleForm.unit = props.goodsOneData.goods.goods_unit
        ruleForm.cost_price = props.goodsOneData.goods.goods_cost_price
        ruleForm.warehouse = props.goodsOneData.goods.goods_warehouse_id
        ruleForm.supply_price = props.goodsOneData.goods.goods_supply_price
        ruleForm.sale_number = props.goodsOneData.goods.goods_sale_number
        ruleForm.return_number = props.goodsOneData.goods.goods_return_number
        ruleForm.link_return_number = props.goodsOneData.goods.goods_link_return_number
        goodsId.value = props.goodsOneData.goods.id
    }
})

// 获取所有仓库
let getWarehouses = () => {
    getWarehouseList().then(r => {
        warehouse.value = r.data.data
    })
}

let returnPage = () => {
    ElMessageBox.confirm('确定返回上一页吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        emits("isShow", true)
    }).catch(() => {

    })
}

const dialogClost = () => {
    dialogVisible.value = false
}

// 创建
const dialogCreate = () => {
    saveWarehouse(warehouseInput.value).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("添加成功")
        }
    })
    dialogVisible.value = false
    getWarehouses()
}

const rules = reactive<FormRules>({
    name: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ],
    date: [
        {
            required: true,
            message: '请选择销售日期',
            trigger: 'change',
        },
    ],
    unit: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ],
    cost_price: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ],
    warehouse: [
        {
            required: true,
            message: '请选择仓库',
            trigger: 'change',
        },
    ],
    supply_price: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ],
    sale_number: [
        { required: false, message: '数据不能为空', trigger: 'blur' },
    ],
    return_number: [
        { required: false, message: '数据不能为空', trigger: 'blur' },
    ],
    link_return_number: [
        { required: false, message: '数据不能为空', trigger: 'blur' },
    ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            let params: SaveGoodsParams = {
                warehouse_id: Number(ruleForm.warehouse),
                cost_price: Number(ruleForm.cost_price),
                supply_price: Number(ruleForm.supply_price),
                sale_number: Number(ruleForm.sale_number),
                return_number: Number(ruleForm.return_number),
                link_return_number: Number(ruleForm.link_return_number),
                name: ruleForm.name,
                date: dateFormat(ruleForm.date),
                unit: ruleForm.unit,
                id: goodsId.value
            }
            if (goodsId.value == 0) {
                saveGoods(params).then(r => {
                    if (r.data.code == 0) {
                        ElMessage.success("创建成功")
                    }
                })
            } else {
                console.log(ruleForm.warehouse)
                modifyGoods(params).then(r => {
                    if (r.data.code == 0) {
                        ElMessage.success("修改成功")
                    }
                })
            }
        } else {
            console.log('error submit!', fields)
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>


<style scoped lang='scss'>
.el-main {
    padding-top: 0;
}

.zt_commit {
    position: absolute;
    right: 60px;
}

.z_a {
    margin-left: 10px;
    color: #aaadb4;
}
</style>