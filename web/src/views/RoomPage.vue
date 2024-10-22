<script setup lang="ts">
import {onMounted} from 'vue'
import {useRoute} from 'vue-router'
import {WSClient} from '@/ws-client'
import type {Option} from '@/ws-client/data-types'
import {layer} from '@layui/layui-vue'


//------------------------------------------------out data--------------------------------------------------------------
const route = useRoute()

onMounted(() => {
  console.log('--->', route.params.id);
  const client = new WSClient({
    url: `${window._globalConfig.roomUrl}?roomId=${route.params.id}`,
    error: (err: any) => {
      layer.msg(err, {time: 1000, icon: 2})
    }
  })
})

</script>

<template>

</template>

<style scoped>

</style>