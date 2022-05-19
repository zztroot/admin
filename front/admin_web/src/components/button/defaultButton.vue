<template>
    <el-button :type="types" style="margin-bottom: 2px;">{{ name }}</el-button>
</template>


<script setup lang='ts'>
defineProps(["name", "types"])
</script>


<style scoped lang='scss'>
</style>