<template>
    <el-card class="box-card zt-card">
        <template #header>
            <div class="card-header">
                <span>基本信息</span>
            </div>
        </template>
        <span>
            <el-input v-model="input" placeholder="请输入角色名称" />
        </span>
    </el-card>
</template>


<script setup lang='ts'>
import { ref, computed, watch, onMounted } from "vue"
import { useStore } from "@/store/index"

let store = useStore()
let input = computed(() => {
    return props.roleInfo.role_name
})
let isCommit = computed(() => {
    return store.getters["getIsCommit"]
})
let emits = defineEmits(["nameData"])
let props = defineProps<{
    roleInfo: any
}>()

// watch 三个参数=>监听者，监听前后方法，配置（immediate立即执行）
watch(
    isCommit, (a, b) => {
        if (a != 0) {
            // 将数据返回给父组件
            emits("nameData", input.value)
        }
    }
)

</script>


<style scoped lang='scss'>
.zt-card {
    margin-bottom: 30px;
}
</style>