<script setup lang="ts">
import {reactive, ref, onMounted} from 'vue'
import API from '@/api'
import {layer} from '@layui/layui-vue'
import {useRouter} from 'vue-router'

//------------------------------------------------------out-------------------------------------------------------------
const router = useRouter()


//-------------------------------------------------------toolbar--------------------------------------------------------
const searchKeys = ref({name: ''})
const onSearch = async () => tableData.value = await API.Room.all(searchKeys.value)
const onAdd = () => {
  layer.prompt({
    title: '输入房间名称',
    formType: 'text',
    value: '',
    btn: [
      {
        text: "创建",
        async callback(layId, data) {
          await API.Room.add({name: data})
          await onSearch()
          layer.close(layId)
        }
      }
    ]
  })
}
const onJoin = (roomId: string) => router.push(`/room/${roomId}`)
const onDel = (roomId: string) => {
  layer.confirm('确定删除此房间？',
      {
        btn: [{
          text: '删除',
          callback: async (id) => {
            await API.Room.del({id: roomId})
            await onSearch()
            layer.close(id)
          }
        },
          {
            text: '取消',
            callback: (id: any) => layer.close(id)
          }]
      }
  )
}

//--------------------------------------------------------table---------------------------------------------------------
const tableColumns = [
  {title: "名称", width: "50%", key: "name"},
  {title: "创建时间", width: "30%", key: "created_at"},
  {title: "在线人数", width: "10%", key: "online_num"},
  {title: "操作", width: "10%", customSlot: "operator", key: "operator", fixed: "right", ignoreExport: true}
]
const tableData = ref([])

onMounted(() => onSearch())
</script>

<template>
  <lay-row class="page-toolbar">
    <lay-col md="12">
      <lay-space size="md">
        <lay-input v-model="searchKeys.name" prefix-icon="layui-icon-search"/>
        <lay-button type="primary">搜索</lay-button>
      </lay-space>
    </lay-col>
    <lay-col md="12" class="right">
      <lay-button type="primary" @click="onAdd">
        <lay-icon type="layui-icon-add-one"></lay-icon>
        创建房间
      </lay-button>
    </lay-col>
  </lay-row>
  <div class="page-body">
    <lay-table :columns="tableColumns" :data-source="tableData" @change="onSearch">
      <template v-slot:operator="{ row }">
        <lay-button size="xs" type="primary" @click="onJoin(row.id)">加入</lay-button>
        <lay-button size="xs" @click="onDel(row.id)">删除</lay-button>
      </template>
    </lay-table>
  </div>
</template>

<style lang="less" scoped>
.page-toolbar {
  background-color: #fff;
  margin-bottom: 16px;
  border-radius: var(--border-radius);
  padding: 16px;

  .right {
    text-align: right;
  }
}

.page-body {
  background-color: #fff;
  border-radius: var(--border-radius);
  padding: 16px;
}
</style>