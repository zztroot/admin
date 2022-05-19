<template>
    <!-- 修改信息 -->
    <el-drawer
        v-model="drawer"
        :title="title"
        :with-header="true"
        :show-close="true"
        :before-close="handleClose"
    >
        <span>
            <slot naem="role_modify"></slot>
            <slot name="role_aside"></slot>
            <el-button class="zt_drawer_button" type="primary" @click="commit()">提交</el-button>
        </span>
    </el-drawer>
</template>


<script setup lang='ts'>
import { ref, computed } from 'vue'
import { ElMessageBox } from 'element-plus'
import { useStore } from '@/store/index'

defineProps(["title"])
let emits = defineEmits(["commit"])

let commit = () => {
    emits("commit", true)
}

const store = useStore()
const drawer = computed(() => {
    return store.getters['getRoleDrawer']
})

const handleClose = (done: () => void) => {
    ElMessageBox.confirm('确定关闭吗？', {
        confirmButtonText: '关闭',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            // 关闭drawer
            store.commit('setRoleDrawer', false)
            done()
        })
        .catch(() => {
            // 关闭失败
        })
}

</script>


<style scoped lang='scss'>
.zt_drawer_button {
    float: right;
    margin-top: 25px;
}
</style>