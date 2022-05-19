<template>这是首页</template>


<script setup lang='ts'>
import { onMounted } from "@vue/runtime-core";
import { httpHeader } from "@/api/api"
import { store } from "@/store/index"

onMounted(() => {
    httpHeader("get", "/test", "", {})
})
</script>


<style scoped lang='scss'>
</style>