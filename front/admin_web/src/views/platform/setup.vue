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
                        <el-form-item label="商品编号" prop="goods">
                            <el-input v-model="ruleForm.goods" />
                        </el-form-item>

                        <el-form-item label="平台售价" prop="price">
                            <el-input v-model="ruleForm.price" />
                        </el-form-item>
                        <el-form-item label="选择平台" prop="platform">
                            <el-select v-model="ruleForm.platform" placeholder="选择平台">
                                <el-option
                                    v-for="val in platfrom"
                                    :label="val.platform_name"
                                    :value="val.id"
                                />
                            </el-select>
                            <span class="z_a">
                                没有平台？
                                <el-button type="text" @click="dialogVisible = true">去添加</el-button>
                            </span>
                        </el-form-item>
                    </el-aside>
                    <el-main>
                        <el-form-item label="平台备注" prop="supply_price">
                            <el-input v-model="ruleForm.comments" />
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
import type { platformStatisticsSetup } from "@/api/platformStatistics"
import { dateFormat } from "@/util/time"
import { savePlatformStatistics, modifyPlatformStatistics } from "@/api/platformStatistics"
import { getPlatfromList, savePlatfrom } from "@/api/platfrom"

const dialogVisible = ref(false)
const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const warehouseInput = ref("")
let ruleForm = reactive({
    platform: '',
    goods: "",
    comments: "",
    price: 0,
})
let platfrom = ref()
let goodsId = ref(0)

let emits = defineEmits(["isShow"])

let props = defineProps(["goodsOneData"])

onMounted(() => {
    getWarehouses()
    if (props.goodsOneData != undefined) {
        ruleForm.platform = props.goodsOneData.plat_form.platform_model.platform_id
        ruleForm.goods = props.goodsOneData.goods.goods_model.goods_code
        ruleForm.comments = props.goodsOneData.plat_form.platform_model.platform_comments
        ruleForm.price = props.goodsOneData.plat_form.platform_model.platform_price
        goodsId.value = props.goodsOneData.plat_form.platform_model.id
    }
})

// 获取所有仓库
let getWarehouses = () => {
    getPlatfromList().then(r => {
        platfrom.value = r.data.data
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
    savePlatfrom(warehouseInput.value).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("添加成功")
        }
    })
    dialogVisible.value = false
    getWarehouses()
}

const rules = reactive<FormRules>({
    price: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ],
    platform: [
        {
            required: true,
            message: '请选择平台',
            trigger: 'change',
        },
    ],
    goods: [
        { required: true, message: '数据不能为空', trigger: 'blur' },
    ]
})

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            let params: platformStatisticsSetup = {
                id: goodsId.value,
                platform_id: Number(ruleForm.platform),
                price: Number(ruleForm.price),
                goods: ruleForm.goods,
                comments: ruleForm.comments
            }
            if (goodsId.value == 0) {
                savePlatformStatistics(params).then(r => {
                    if (r.data.code == 0) {
                        ElMessage.success("创建成功")
                    }
                })
            } else {
                modifyPlatformStatistics(params).then(r => {
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