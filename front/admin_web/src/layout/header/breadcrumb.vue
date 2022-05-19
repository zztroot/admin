<template>
    <el-breadcrumb separator="/">
        <el-breadcrumb-item v-for="item in breadcrumbList">
            <a :href="item.path" v-if="item.children.length < 1">{{ item.meta.title }}</a>
            <span v-else>{{ item.meta.title }}</span>
        </el-breadcrumb-item>
    </el-breadcrumb>
</template>

<script setup lang='ts'>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

// 面包屑导航功能
const route = useRoute()
const breadcrumbList: any = ref([])

const caleBreadcrumb = () => {
    breadcrumbList.value = route.matched.filter(
        (item) => item.meta && item.meta.title
    )
}

watch(
    () => route.path, // 也可以直接写 route ; 即  () => route.path  替换成 route
    () => {
        caleBreadcrumb()
    },
    {
        immediate: true
    }
)
</script>


<style scoped lang='scss'>
</style>