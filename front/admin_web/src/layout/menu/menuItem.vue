<template>
    <template v-for="menu in itemList" :key="menu.path">
        <el-sub-menu v-if="menu.children && menu.children.length > 0" :index="menu.path">
            <template #title>
                <el-icon>
                    <component :is="menu.icon"></component>
                </el-icon>
                <span>{{ menu.title }}</span>
            </template>
            <menu-item :itemList="menu.children"></menu-item>
        </el-sub-menu>
        <el-menu-item style="color: #f4f4f5;" v-else :index="menu.path">
            <el-icon>
                <component :is="menu.icon"></component>
            </el-icon>
            <template #title>
                <span>{{ menu.title }}</span>
            </template>
        </el-menu-item>
    </template>
</template>  

<script lang="ts" setup>
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { store } from '@/store/index'

// 接受父组件数据
defineProps(["itemList"])

const route = useRoute()
const isCollapse = computed(() => {
    return store.getters['getCollapse']
})

const activeIndex = computed(() => {
    const { path } = route
    return path
})
console.log(activeIndex.value)

</script>

<style lang="scss" scoped>
</style>