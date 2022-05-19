<template>
    <el-row :gutter="20">
        <el-col :span="5">
            <div class="grid-content bg-purple">
                <el-input v-model="userNameInput" placeholder="用户名搜索">
                    <template #prepend>名称</template>
                </el-input>
            </div>
        </el-col>
        <el-col :span="5">
            <div class="grid-content bg-purple">
                <el-autocomplete
                    v-model="sexInput"
                    :fetch-suggestions="querySearch"
                    clearable
                    class="inline-input"
                    placeholder="性别搜索"
                    @select="handleQuerySelect"
                    type="text"
                >
                    <template #prepend>性别</template>
                </el-autocomplete>
            </div>
        </el-col>
        <DefaultButton name="重置" @click="reset"></DefaultButton>
        <DefaultButton name="搜索" @click="search" types="primary"></DefaultButton>
        <div class="createUser">
            <DefaultButton name="创建" @click="dialogVisible = true" types="primary"></DefaultButton>
        </div>
        <!-- 新建用户弹窗 -->
        <BaseDialog
            :dialogVisible="dialogVisible"
            @dialogClost="dialogClost"
            @dialogCreate="dialogCreate"
            title="创建用户"
            buttonName="创建"
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
                                v-model="inputForm.password"
                                placeholder="请输入密码"
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
    </el-row>
</template>


<script setup lang='ts'>
import { ref, reactive, onMounted, computed } from "vue"
import DefaultButton from '@/components/button/defaultButton.vue'
import BaseDialog from "@/components/dialog/baseDialog.vue"
import { getRoleList } from "@/api/role"
import type { httpRequests } from "@/api/api"
import { ElMessage } from "element-plus"
import { create, getlist } from "@/api/user"
import { useStore } from "@/store/index"
import type { UserSearch } from "@/store/index"

let store = useStore()
let userNameInput = ref('')
let dialogVisible = ref(false)
let querySexId = ref(0)
let sexId = ref()
let queryRoleId = ref()
let sexInput = ref()
let sexMap = new Map(
    [
        ["男", 1],
        ["女", 2]
    ]
)

let props = defineProps<{
    roleList?: any
}>()

let emits = defineEmits(["handlerSearch"])

let search = () => {
    let searchUser: UserSearch = {
        userName: userNameInput.value,
        sexId: sexMap.get(sexInput.value),
        isSearch: store.getters["getUserHandlerSearch"].isSearch + 1
    }
    store.commit("setUserHandlerSearch", searchUser)
}

let reset = () => {
    userNameInput.value = ""
    sexInput.value = ""
    let searchUser: UserSearch = {
        userName: userNameInput.value,
        sexId: sexMap.get(sexInput.value),
        isSearch: store.getters["getUserHandlerSearch"].isSearch + 1
    }
    store.commit("setUserHandlerSearch", searchUser)
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
let inputForm = reactive({
    user_name: '',
    password: '',
    again_password: '',
    real_name: '',
    phone: '',
    sex: '',
    age: '',
    role: ''
})

let rules = reactive({
    user_name: [
        { required: true, message: '账号不能为空', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '密码不能为空', trigger: 'blur' }
    ],
    again_password: [
        { required: true, message: '密码不能为空', trigger: 'blur' }
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

const handleQuerySelect = (item: any) => {
    querySexId.value = item.id
}

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

// 创建用户
const dialogCreate = () => {
    if (queryRoleId.value == -1) {
        ElMessage.warning("请选择角色")
        return
    }
    if (inputForm.password != inputForm.again_password) {
        ElMessage.warning("两次密码不一致")
        return
    }
    create(inputForm.user_name, inputForm.password, inputForm.real_name, inputForm.phone, Number(inputForm.age), sexId.value, queryRoleId.value, inputForm.again_password).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("创建成功")
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
    inputForm.password = ""
    inputForm.again_password = ""
    sexId.value = ref()
    queryRoleId.value = ref()
}

</script>


<style scoped lang='scss'>
:deep(.el-autocomplete) {
    width: 100%;
}

.createUser {
    position: absolute;
    right: 0px;
}
</style>