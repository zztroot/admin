<template>
    <el-container class="layout">
        <el-aside class="aside" width="auto">
            <Menu></Menu>
        </el-aside>
        <el-container>
            <el-header class="header">
                <Header></Header>
            </el-header>
            <el-main class="main">
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script setup lang="ts">
import Header from '@/layout/header/header.vue'
import Menu from '@/layout/menu/menu.vue'

</script>

<style lang="scss" scoped>
.layout {
    height: 100%;
    .aside {
        background-color: #304156;
    }
    .header {
        height: 50px;
        border-bottom: 1px solid #e5e5e5;
        display: flex;
        align-items: center;
    }
    .main {
    }
}
</style>