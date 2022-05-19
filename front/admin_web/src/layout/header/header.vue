<template>
    <Collapse></Collapse>
    <Breadcrumb class="breadcrumb"></Breadcrumb>
    <Avatar></Avatar>
</template>

<script lang="ts" setup>
import Collapse from './collapse.vue'
import Breadcrumb from './breadcrumb.vue'
import Avatar from './avatar.vue'
</script> 

<style lang="scss" scoped>
.breadcrumb {
    margin-left: 15px;
}
</style>