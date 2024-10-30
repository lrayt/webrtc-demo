<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {useRoute} from 'vue-router'
import {WSClient} from '@/ws-client'
import {LogLevel, type LogInfo} from '@/common/logger'
import {type UserInfo} from '@/common/user'
import API from '@/api'
import {type Message, Actions} from "@/ws-client/data-types";
import {NowDate} from "@/common/date";


//------------------------------------------------out data--------------------------------------------------------------
const route = useRoute()
const inputText = ref<string>('')
const userList = ref<UserInfo[]>([])
const msgList = ref<Message[]>([])
const logList = ref<LogInfo[]>([])

let client: WSClient

const onLoadUser = async () => {
  userList.value = await API.Room.users(route.params.id)
  console.log('--->', userList.value)
}
const onSend = () => {
  console.log('--->', inputText.value);
  const msg: Message = {
    uid: client.getUID(),
    action: Actions.Message,
    content: inputText.value,
    date: NowDate()
  }
  msgList.value.push(msg)
  client.send(msg)
}


onMounted(() => {
  client = new WSClient({
    url: `${window._globalConfig.roomUrl}?roomId=${route.params.id}`,
    onLog: (logInfo: LogInfo) => logList.value.push(logInfo),
    onUserJoin: async () => onLoadUser()
  })
})

</script>

<template>
  <div class="main-container">
    <lay-row space="8" class="row">
      <lay-col md="4" class="col">
        <ul class="grid">
          <li v-for="o in userList">{{ o.id }}</li>
        </ul>
      </lay-col>
      <lay-col md="14" class="col">
        <div class="grid"></div>
      </lay-col>
      <lay-col md="6" class="col">
        <div class="grid">
          <ul class="msg-box">
            <li v-for="o in msgList">
              <p>{{ o.uid }}</p>
              <p>{{ o.content }}</p>
            </li>
          </ul>
          <div class="input-box">
            <lay-textarea placeholder="请输入描述" v-model="inputText" @keyup:enter.native="onSend"/>
          </div>
        </div>
      </lay-col>
    </lay-row>
  </div>
  <div class="log-container">
    <lay-timeline class="box">
      <lay-timeline-item v-for="o in logList">
        <template #dot>
          <lay-icon v-if="o.level===LogLevel.Info" type="layui-icon-face-smile" color="green"/>
          <lay-icon v-else-if="o.level===LogLevel.Warning" type="layui-icon-face-surprised" color="yellow"/>
          <lay-icon v-else-if="o.level===LogLevel.Error" type="layui-icon-face-cry" color="red"/>
          <lay-icon v-else type="layui-icon-notice" color="#009688"/>
        </template>
        <template #title>
          {{ o.date }}
        </template>
        <p>
          {{ o.content }}
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
      background-color: #ffffff;
      padding: 16px;

      li {
        padding-left: 16px;
        height: 32px;
        line-height: 32px;
        border: 1px solid var(--border-color);
        border-radius: var(--border-radius);
        margin-bottom: 8px;
      }

      .msg-box {
        width: 100%;
        height: calc(100% - 116px);
        border: 1px solid var(--border-color);
        border-radius: var(--border-radius);
        margin-bottom: 16px;
      }

      .input-box {
        width: 100%;
        height: 100px;
      }
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