<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {useRoute} from 'vue-router'
import {WSClient} from '@/ws-client'
import type {Message} from '@/ws-client/data-types'
import {layer} from '@layui/layui-vue'


//------------------------------------------------out data--------------------------------------------------------------
const route = useRoute()

const userList = ref()
const msgList = ref()
const logList = ref<Message[]>([])


onMounted(() => {
  console.log('--->', route.params.id);
  const client = new WSClient({
    url: `${window._globalConfig.roomUrl}?roomId=${route.params.id}`,
    log: (msg: Message) => {
      console.log('--->', msg);
      logList.value.push(msg)
    },
    error: (err: any) => {
      layer.msg(err, {time: 1000, icon: 2})
    }
  })
})

</script>

<template>
  <div class="main-container">
    <lay-row space="8" class="row">
      <lay-col md="4" class="col">
        <div class="grid"></div>
      </lay-col>
      <lay-col md="14" class="col">
        <div class="grid"></div>
      </lay-col>
      <lay-col md="6" class="col">
        <div class="grid"></div>
      </lay-col>
    </lay-row>
  </div>
  <div class="log-container">
    <lay-timeline class="box">
      <lay-timeline-item v-for="o in logList">
        <template #dot>
          <lay-icon type="layui-icon-face-smile" color="red"></lay-icon>
        </template>
        <template #title>
          {{ o.date }}
        </template>
        <p>
          {{ o.action }}
        </p>
      </lay-timeline-item>
    </lay-timeline>
  </div>
</template>

<style lang="less" scoped>
@log-height: 300px;

.main-container {
  width: 100%;
  height: calc(100% - @log-height);
  background-color: #f2f2f2;
  padding: 8px 0 0 8px;

  .row {
    width: 100%;
    height: 100%;
  }

  .col {
    height: 100%;

    .grid {
      height: 100%;
      border: 1px solid var(--border-color);
      border-radius: var(--border-radius);
    }
  }
}

.log-container {
  width: 100%;
  height: @log-height;
  background-color: var(--bg-color);
  padding: 0 8px 8px 8px;

  .box {
    padding: 16px;
    height: 100%;
    width: 100%;
    border: 1px solid var(--border-color);
    border-radius: var(--border-radius);
    overflow-y: auto;
    background-color: #ffffff;
  }
}
</style>