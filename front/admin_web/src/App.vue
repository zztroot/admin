
<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { watch, computed } from "vue"
import { useRoute } from 'vue-router'
import { useStore } from "@/store/index"
import type { Setup } from "@/store/index"

let route = useRoute()
let store = useStore()
let path = computed(() => {
  return route.path
})

watch(path, (a, b) => {
  if (a == "/goods/setup") {
    let setup: Setup = {
      goodsSetup: true
    }
    store.commit("setSetupShow", setup)
  }
  if (a === "/goods") {
    let setup: Setup = {
      goodsSetup: false
    }
    store.commit("setSetupShow", setup)
  }
  if (a === "/platform") {
    let setup: Setup = {
      platformSetup: false
    }
    store.commit("setSetupShow", setup)
  }
  if (a == "/platform/platformSetup") {
    let setup: Setup = {
      platformSetup: true
    }
    store.commit("setSetupShow", setup)
  }
}, { "immediate": true })

</script>


<style>
</style>
