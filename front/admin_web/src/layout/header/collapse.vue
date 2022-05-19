<template>
    <el-icon @click="changeIcon" class="icons">
        <component :is="status ? Expand : Fold"></component>
    </el-icon>
</template>

<script setup lang='ts'>
import { store } from '@/store/index'
import { Fold, Expand } from '@element-plus/icons-vue'
import { ref, computed } from 'vue'

const status = computed(() => {
    return store.getters['getCollapse']
})

// 图标切换函数
const changeIcon = () => {
    store.commit('setCollapse', !status.value)
}
</script>


<style scoped lang='scss'>
.icons {
    color: #303133;
    font-size: 20px;
    align-items: center;
    cursor: pointer;
}
</style>