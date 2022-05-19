<template>
    <el-container>
        <el-main>
            <DefaultTable :tableData="tableData" @roleData="roleData" @roleCreate="roleCreate"></DefaultTable>
        </el-main>
        <el-aside width="30%">
            <el-card class="box-card" shadow="never">
                <!-- 修改角色 -->
                <RoleModify slot="role_modify" :roleInfo="roleInfo" @nameData="getName"></RoleModify>
                <RoleSetAuth
                    slot="role_aside"
                    :data="authData"
                    :active="actives"
                    @authData="getAuthData"
                ></RoleSetAuth>
                <el-button class="zt_drawer_button" type="primary" @click="commit()">提交</el-button>
            </el-card>
        </el-aside>
    </el-container>
</template>

<script setup lang='ts'>
import RoleSetAuth from './roleSetAuth.vue'
import RoleModify from './roleModify.vue'
import DefaultTable from './roleTable.vue'
import Pagination from '@/components/pagination/default.vue'
import { onMounted, reactive, ref } from 'vue'
import {
    deleteRole,
    getRoleList,
    getMenuList,
    modifyRole
} from "@/api/role"
import type { httpRequests } from "@/api/api"
import { useStore } from "@/store/index"
import { ElMessage } from 'element-plus'

const tableData = ref([])
let authData = ref([])
let roleInfo = ref([])
let actives = ref([])
let store = useStore()
let authStr = ref("")
let nameStr = ref("")
let roleId = ref(0)

onMounted(() => {
    // 获取所有角色
    let req: httpRequests = {
        params: {},
        headers: {}
    }
    // 获取全部菜单
    getMenuList(req).then(response => {
        authData.value = response.data.data
    })
    getList()
})

// 查询所有角色
let getList = () => {
    // 获取所有角色
    let req: httpRequests = {
        params: {},
        headers: {}
    }
    getRoleList(req).then(response => {
        tableData.value = response.data.data
    })
}

// 提交方法
let commit = () => {
    store.commit("setIsCommit", store.getters["getIsCommit"] + 1)
}

// 选择的角色回调方法
let roleData = (data: any) => {
    roleInfo.value = data
    let { id } = data
    roleId.value = id
    store.commit("setAuth", JSON.parse(data.role_extend))
    store.commit("setIsSelect", store.getters["getIsSelect"] + 1)
}

// 获取到name
let getName = (name: string) => {
    if (name == undefined) {
        ElMessage.error("请选择需要修改的角色")
        return
    }
    nameStr.value = name
    setTimeout(() => {
        if (nameStr.value !== "" && authStr.value != "") {
            setUp(roleId.value)
        }
    }, 200);
}

// 获取到权限数据
let getAuthData = (data: any) => {
    let auth = JSON.stringify(data)
    authStr.value = auth
}

// 保存
let setUp = (id: number) => {
    let req: httpRequests = {
        params: {
            "id": id,
            "role_name": nameStr.value,
            "role_extend": authStr.value
        },
        headers: {}
    }
    modifyRole(req).then(r => {
        if (r.data.code == 0) {
            ElMessage.success("提交成功")
            getList()
        }
    })
}

// 创建成功后，再次查询
let roleCreate = () => {
    getList()
}

</script>


<style scoped lang='scss'>
.zt_drawer_button {
    float: right;
    margin: 25px 0;
}
</style>