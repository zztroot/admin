<template>
    <div class="content">
        <div class="login">
            <el-form ref="ruleFormRef" :model="inputForm" :rules="rules">
                <div class="suffix" style="font-size: 23px;color: rgb(252, 251, 250);">系统管理登录</div>
                <div class="suffix">
                    <el-form-item prop="name">
                        <el-input placeholder="请输入账号" v-model="inputForm.name" prefix-icon="User"></el-input>
                    </el-form-item>
                </div>
                <div class="suffix">
                    <el-form-item prop="password">
                        <el-input
                            placeholder="请输入密码"
                            v-model="inputForm.password"
                            type="password"
                            prefix-icon="Lock"
                            v-if="!pwd_type"
                        >
                            <template #suffix>
                                <el-icon class="zt_view_icon" @click="showPwd()">
                                    <component is="View"></component>
                                </el-icon>
                            </template>
                        </el-input>
                        <el-input
                            placeholder="请输入密码"
                            v-model="inputForm.password"
                            type="text"
                            prefix-icon="Lock"
                            v-else
                        >
                            <template #suffix>
                                <el-icon class="zt_view_icon" @click="showPwd()">
                                    <component is="View"></component>
                                </el-icon>
                            </template>
                        </el-input>
                    </el-form-item>
                </div>
            </el-form>
            <div class="suffix">
                <el-button
                    type="primary"
                    style="width: 100%"
                    @click="login(router)"
                    v-on:keyup.enter="login(router)"
                >登录</el-button>
            </div>
        </div>
    </div>
</template>


<script setup lang='ts'>
import { onMounted, reactive, ref, watch } from 'vue'
import { store } from '@/store/index'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from "element-plus"
import { http } from "@/api/api"

const router = useRouter()
const route = useRoute()
const reply = ref()
const pwd_type = ref(false)

let inputForm = reactive({
    name: '',
    password: ''
})

let rules = reactive({
    name: [
        { required: true, message: '账号不能为空', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '密码不能为空', trigger: 'blur' }
    ]
})



let login = (r: any) => {
    if (inputForm.name === "" || inputForm.password === "") {
        ElMessage.error("账号或密码不能为空")
        return
    }
    let params = {
        "user_name": inputForm.name,
        "user_password": inputForm.password
    }
    http("post", '/user/login', params).then(response => {
        if (response.data.code !== 0) {
            ElMessage.error(response.data.message)
            return
        }
        if (response.data.data.token === "") {
            return
        }
        localStorage.setItem("user", JSON.stringify(response.data.data))
        store.commit("setToken", response.data.data.token)
        ElMessage.success("欢迎回来")
        setTimeout(() => {
            r.push({
                path: '/home'
            })
        }, 1000);
    })
}

let showPwd = () => {
    pwd_type.value = !pwd_type.value
}

// 监听回车事件
document.onkeydown = (e) => {
    if (e.keyCode == 13 && route.path == "/") {
        login(router)
    }
}

</script>


<style scoped lang='scss'>
.content {
    height: 100%;
    width: 100%;
    background-color: #304156;
    /* padding-top: 100px; */
    overflow: hidden;
    .login {
        width: 350px;
        height: 300px;
        margin-top: 150px;
        margin-left: auto;
        margin-right: auto;
        text-align: center;
        padding: 10px;
        border-radius: 50%;
        .suffix {
            margin-top: 30px;
            .zt_view_icon {
                cursor: pointer;
                display: flex;
                height: inherit;
            }
            .zt_view_icon:hover {
                color: #304156;
            }
        }
    }
}
</style>