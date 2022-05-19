<template>
    <el-container>
        <el-header>
            <!-- 搜索框 -->
            <UserHeader :roleList="roleList"></UserHeader>
        </el-header>
        <el-main>
            <!-- 主题内容 -->
            <UserMain :roleList="roleList"></UserMain>
        </el-main>
        <el-footer>
            <!-- 尾部 -->
        </el-footer>
    </el-container>
</template>


<script setup lang='ts'>
import UserMain from "./user/userMain.vue"
import UserHeader from "./user/userHeader.vue"
import { ref, onMounted } from "vue"
import type { httpRequests } from "@/api/api"
import { getRoleList } from "@/api/role"

let roleList = new Array

// 进入页面执行
onMounted(() => {
    // 获取角色列表
    // 获取所有角色
    let req: httpRequests = {
        params: {},
        headers: {}
    }
    getRoleList(req).then(response => {
        response.data.data.forEach((e: any) => {
            let role = {
                "value": e.role_name,
                "id": e.id
            }
            roleList.push(role)
        });
    })
})

</script>


<style scoped lang='scss'>
</style>