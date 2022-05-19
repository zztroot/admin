<template>
    <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column label="ID" width="80">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.$index + 1 }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="用户名" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_real_name }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="账号" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_name }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="手机号" width="200">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_phone }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="年龄" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_age }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="性别" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_sex }}</span>
                </div>
            </template>
        </el-table-column>
        <el-table-column label="角色" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                    <span>{{ scope.row.user_role }}</span>
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
                <el-button size="small" type="primary" @click="handleEdit(scope.row)">修改</el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
    <!-- 修改用户弹窗 -->
    <BaseDialog
        :dialogVisible="dialogVisible"
        @dialogClost="dialogClost"
        @dialogCreate="dialogCreate"
        title="修改用户"
        buttonName="修改"
        width="20%"
    >
        <template #dialogMain>
            <el-form :model="inputForm" :rules="rules">
                <el-space direction="horizontal" :fill="true">
                    <el-form-item prop="user_name">
                        <el-input v-model="inputForm.user_name" placeholder="请输入账号"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input
                            v-model="inputForm.new_password"
                            placeholder="请输入新密码"
                            type="password"
                        ></el-input>
                    </el-form-item>
                    <el-form-item prop="again_password">
                        <el-input
                            v-model="inputForm.again_password"
                            placeholder="请再次输入密码"
                            type="password"
                        ></el-input>
                    </el-form-item>
                    <el-form-item prop="real_name">
                        <el-input v-model="inputForm.real_name" placeholder="请输入用户真实名称"></el-input>
                    </el-form-item>
                    <el-form-item prop="phone">
                        <el-input v-model="inputForm.phone" placeholder="请输入用户手机号"></el-input>
                    </el-form-item>
                    <el-form-item prop="age">
                        <el-input v-model="inputForm.age" placeholder="请输入用户年龄"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-autocomplete
                            v-model="inputForm.sex"
                            :fetch-suggestions="querySearch"
                            clearable
                            class="inline-input"
                            placeholder="请选择性别"
                            @select="handleSexSelect"
                            type="text"
                        ></el-autocomplete>
                    </el-form-item>
                    <el-autocomplete
                        v-model="inputForm.role"
                        :fetch-suggestions="queryRoleList"
                        clearable
                        class="inline-input"
                        placeholder="请选择角色"
                        @select="handleCreateSelect"
                        type="text"
                    ></el-autocomplete>
                </el-space>
            </el-form>
        </template>
    </BaseDialog>
</template>


<script setup lang='ts'>
import DefaultButton from '@/components/button/defaultButton.vue'
import BaseDialog from "@/components/dialog/baseDialog.vue"
import { ref, reactive, onMounted } from "vue"
import { deleteUserOne, modify, getUserOne } from "@/api/user"
import { ElMessage, ElMessageBox } from "element-plus"
import { useStore } from "@/store/index"

let props = defineProps<{
    roleList?: any,
    tableData?: any,
    loading?: any
}>()

const dialogVisible = ref(false)
let userName = ref("")
let fill = ref(true)
const size = ref(30)
let store = useStore()
let querySexId = ref()
let sexId = ref(0)
let queryRoleId = ref(0)
let id = ref()
let inputForm = reactive({
    user_name: '',
    again_password: '',
    real_name: '',
    phone: '',
    sex: '',
    age: '',
    role: '',
    new_password: '',
    role_id: 0
})
let sexMap = new Map(
    [
        ["男", 1],
        ["女", 2]
    ]
)

// 用户修改
const handleEdit = (item: any) => {
    dialogVisible.value = true
    inputForm.user_name = item.user_name
    inputForm.real_name = item.user_real_name
    inputForm.phone = item.user_phone
    inputForm.sex = item.user_sex
    inputForm.age = item.user_age
    inputForm.role = item.user_role
    inputForm.role_id = item.user_role_id
    id.value = item.id
}

// 用户删除
const handleDelete = (index: number) => {
    ElMessageBox.confirm('确定删除吗？', {
        "cancelButtonText": "取消",
        "confirmButtonText": "确定"
    }).then(() => {
        deleteUserOne(index).then(r => {
            if (r.data.code == 0) {
                ElMessage.success("删除成功")
                store.commit("setIsQueryUserList", store.getters["getIsQueryUserList"] + 1)
            }
        })
    }).catch(() => {

    })
}

const querySearch = (queryString: string, cb: any) => {
    cb([
        { "value": "男", "id": 1 },
        { "value": "女", "id": 2 },
    ])
}

const queryRoleList = (q: string, cb: any) => {
    cb(props.roleList)
}

// 创建用户参数规则和校验
let rules = reactive({
    user_name: [
        { required: true, message: '账号不能为空', trigger: 'blur' },
    ],
    real_name: [
        { required: true, message: '不能为空', trigger: 'blur' }
    ],
    phone: [
        { required: true, message: '不能为空', trigger: 'blur' }
    ],
    sex: [
        { required: true, message: '不能为空', trigger: 'blur' }
    ],
    age: [
        { required: true, message: '不能为空', trigger: 'blur' }
    ],
})

const handleCreateSelect = (item: any) => {
    queryRoleId.value = item.id
}

const handleSexSelect = (item: any) => {
    sexId.value = item.id
}

// 关闭弹窗
const dialogClost = () => {
    dialogVisible.value = false
}

// 修改用户
const dialogCreate = () => {
    if (id.value == 0) {
        return
    }
    let roleId: number
    let sextempId: number | undefined
    roleId = queryRoleId.value
    sextempId = sexId.value
    if (inputForm.sex != "") {
        let tempId = sexMap.get(inputForm.sex)
        if (tempId != sexId.value) {
            sextempId = tempId
        }
    }
    if (queryRoleId.value == inputForm.role_id) {
        roleId = inputForm.role_id

    }
    if (inputForm.new_password != inputForm.again_password) {
        ElMessage.warning("两次密码不一致")
        return
    }
    modify(inputForm.user_name, "", inputForm.real_name, inputForm.phone, Number(inputForm.age), sextempId, roleId, inputForm.new_password, id.value, inputForm.again_password).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("修改成功")
            queryRoleId.value = -1
            // 通知其他组件获取用户列表
            store.commit("setIsQueryUserList", store.getters["getIsQueryUserList"] + 1)
            dialogVisible.value = false
            clearCreateData()
        } else {
            ElMessage.error(String(r.data.message))
            return
        }
    })
}

let clearCreateData = () => {
    inputForm.user_name = ""
    inputForm.real_name = ""
    inputForm.phone = ""
    inputForm.sex = ""
    inputForm.age = ""
    inputForm.role = ""
    inputForm.role_id = 0
    id.value = 0
    sexId.value = 0
    queryRoleId.value = 0
    id.value = 0
}
</script>

<style scoped lang='scss'>
:deep(.el-autocomplete) {
    width: 100%;
}
</style>