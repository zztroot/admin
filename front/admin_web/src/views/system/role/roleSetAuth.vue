<template>
    <el-card class="box-card">
        <template #header>
            <div class="card-header">
                <span>权限设置</span>
            </div>
        </template>
        <el-tree
            ref="treeRef"
            :data="data"
            show-checkbox
            default-expand-all
            node-key="id"
            highlight-current
            :props="defaultProps"
        />
        <div class="buttons">
            <el-button @click="resetChecked" type="danger">重置</el-button>
        </div>
    </el-card>
</template>

<script setup lang='ts'>
import { computed, nextTick, onMounted, ref, watch } from 'vue'
import type { ElTree } from 'element-plus'
import { useStore } from "@/store/index"

let store = useStore()
let isSelect = computed(() => {
    return store.getters["getIsSelect"]
})
let activeData = computed(() => {
    return store.getters["getAuth"]
})
let isCommit = computed(() => {
    return store.getters["getIsCommit"]
})

let props = defineProps<{
    data?: any,
    active?: any,
}>()

let emits = defineEmits(["authData"])

interface Tree {
    id: number
    label: string
    children?: Tree[]
}

const treeRef = ref<InstanceType<typeof ElTree>>()
const selectData: any = ref()

// watch 三个参数=>监听者，监听前后方法，配置（immediate立即执行）
watch(
    isSelect, (a, b) => {
        if (a != 0) {
            // setTimeout -> 定时器，可能定时器跑完的时候，页面还没渲染
            // nextTick -> 下一次页面渲染完成后，触发
            nextTick(() => {
                treeRef.value!.setCheckedNodes(
                    activeData.value,
                    false
                )
            })
            store.commit("setIsSelect", 0)
        }

    }, { "immediate": true }
)

watch(isCommit, (a, b) => {
    if (a != 0) {
        nextTick(() => {
            emits("authData", treeRef.value!.getCheckedNodes(true, false))
        })
        store.commit("setIsCommit", 0)
    }
}, { "immediate": true })

// 权限重置方法
const resetChecked = () => {
    treeRef.value!.setCheckedKeys([], false)
}

const defaultProps = {
    children: 'children',
    label: 'title',
}

</script>


<style scoped lang='scss'>
.buttons {
    margin: 30px 0 10px 0;
    float: right;
}
</style>