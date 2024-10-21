import Axios from 'axios'
import RoomAPI from '@/api/room'

const client = Axios.create({
    baseURL: window._global_config.serverUrl,
    timeout: 5000,
})

export default {
    Room: new RoomAPI(client)
}