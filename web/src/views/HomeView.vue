<script setup lang="ts">
import {ref, onMounted} from 'vue'
import API from '@/api'
import {Plus} from '@element-plus/icons-vue'

interface RoomInfo {
  Id: string,
  Name: string,
  CreateTime: string,
}

const roomName = ref<string>('')
const roomList = ref<RoomInfo[]>([])
const onLoad = async () => {
  const res = await API.Room.loadAll()
  roomList.value = res.data.data || []
}
const onAdd = async () => {
  await API.Room.add({name: roomName.value})
  await onLoad()
}

const onJoin = async (id: string) => {
  console.log('====>', id);
}

const onDel = async () => {

}


onMounted(() => onLoad())

</script>

<template>
  <div class="container">
    <div class="action-row">
      <el-space wrap>
        <el-input v-model="roomName" style="width: 240px" placeholder="请输入房间名"/>
        <el-button type="primary" :icon="Plus" @click="onAdd">创建房间</el-button>
      </el-space>
    </div>
    <el-table :data="roomList" style="width: 100%">
      <el-table-column prop="name" label="房间名称"/>
      <el-table-column prop="created_at" label="创建时间"/>
      <el-table-column prop="online_num" label="在线人数"/>
      <el-table-column fixed="right" label="操作">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="onJoin(scope.id)">加入</el-button>
          <el-button link type="primary" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<style lang="less" scoped>
.container {
  padding: 32px 16px;
  width: 1280px;
  margin: auto;
}

.action-row {
  text-align: right;
  margin-bottom: 16px;
}

.room {
  height: 54px;
  background-color: red;
}
</style>