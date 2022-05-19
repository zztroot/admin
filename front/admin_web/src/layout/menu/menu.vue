<template>
    <MenuLogo class="layout-logo" v-if="!status"></MenuLogo>
    <el-menu
        :default-active="activeIndex"
        class="el-menu-vertical-demo"
        :collapse="isCollapse"
        @open="handleOpen"
        @close="handleClose"
        background-color="#304156"
        router
    >
        <MenuItem :itemList="MenuItemList"></MenuItem>
    </el-menu>
</template>  

<script lang="ts" setup>
import MenuItem from '@/layout/menu/menuItem.vue'
import MenuLogo from '@/layout/menu/menuLogo.vue'
import { computed, onMounted, reactive, ref } from 'vue'
import { store } from '@/store/index'
import { useRoute } from 'vue-router'

const status = computed(() => {
    return store.getters['getCollapse']
})

// // 接受父组件数据
// defineProps(["itemList"])

const route = useRoute()
const isCollapse = computed(() => {
    return store.getters['getCollapse']
})

const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}

const activeIndex = computed(() => {
    const { path } = route
    return path
})

// 菜单数据，后期从后端获取
let MenuItemList = JSON.parse(String(localStorage.getItem('user'))).user.permissions

</script>

<style lang="scss" scoped>
.el-menu {
    border-right: none;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 230px;
    // min-height: 400px;
}

:deep(.el-sub-menu .el-sub-menu__title) {
    color: #f4f4f5 !important;
}

:deep(.el-menu .el-menu-item) {
    color: #bfcbd9 !important;
}

// .el-menu-item {
//     color: #bfcbd9 !important;
// }

:deep(.el-menu-item.is-active) {
    color: #409eff !important;
}

:deep(.is-opened .el-menu-item) {
    background-color: #1f2d3d !important;
}

:deep(.el-menu-item:hover) {
    background-color: #001528 !important;
}
</style>