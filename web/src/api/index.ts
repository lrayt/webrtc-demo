import Axios from 'axios'
import RoomAPI from '@/api/room-api'
import {layer} from "@layui/layui-vue"

const client = Axios.create({
    baseURL: `${window._globalConfig.serverUrl}/api/v1`,
    timeout: 10000
})

client.interceptors.response.use(({data, status}) => {
    if (status !== 200) {
        layer.msg("网络异常", {time: 1000, icon: 2})
        return Promise.reject('网络异常')
    } else if (data.code !== 10000) {
        layer.msg(data.msg, {time: 1000, icon: 2})
        return Promise.reject(data.msg)
    } else {
        return data.data
    }
}, error => {
    console.error(error)
    return Promise.reject('请求发生错误')
})

export default {
    Room: new RoomAPI(client)
}

