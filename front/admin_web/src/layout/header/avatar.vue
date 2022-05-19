<template>
    <div class="avatar">
        <span style="margin-right: 30px;">
            <a href="https://github.com/zztroot/admin" target="_blank">Github地址</a>
        </span>
        <span>
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
        </span>
        <span class="user_name">
            <el-dropdown @visible-change="isClose">
                <span class="el-dropdown-link">
                    {{ userName }}
                    <el-icon class="el-icon--right">
                        <arrow-down v-if="!dropDown" />
                        <arrow-Up v-if="dropDown" />
                    </el-icon>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>个人信息</el-dropdown-item>
                        <el-dropdown-item @click="outLogin(router)">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </span>
    </div>
</template>


<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import { ElMessage } from "element-plus"
import { useRouter } from 'vue-router'

let router = useRouter()
const dropDown = ref(false)
const isClose = (val: boolean) => {
    dropDown.value = val
}
let userName = ref("")

onMounted(() => {
    let user = JSON.parse(String(localStorage.getItem("user")))
    userName.value = user.user.real_name
})

let outLogin = (r: any) => {
    ElMessage.success("退出成功")
    localStorage.clear()
    setTimeout(() => {
        r.push({
            path: '/'
        })
    }, 1000);
}

</script>


<style scoped lang='scss'>
.avatar {
    right: 50px;
    position: absolute;
    align-items: center;
    display: flex;
    .user_name {
        margin-left: 5px;
        cursor: pointer;
        color: #606266;
    }
}
</style>