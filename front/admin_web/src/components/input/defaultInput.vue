<template>
    <el-input class="input-size" v-model="input1" size="default" :placeholder="name" />
</template>


<script setup lang='ts'>
import { ref } from 'vue'
const props = defineProps({
    name: String
})
const input1 = ref('')
</script>


<style scoped lang='scss'>
.input-size {
    width: 240px;
    margin-right: 30px;
}
</style>