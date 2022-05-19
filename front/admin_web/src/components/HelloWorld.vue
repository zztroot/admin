<template>
  <h1>{{ showCount }}</h1>

  <button type="button" @click="addBth">新增</button>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useStore } from '../store'

defineProps<{ msg: string }>()
const store = useStore()
const count = ref(0)
const showCount = computed(() => {
  return store.getters['getCount']
})

// 函数
const addBth = () => {
  store.commit('setCount', ++count.value)
}
</script>

<style scoped>
</style>
