import type {AxiosInstance} from 'axios'

export default class RoomAPI {
    private readonly client: AxiosInstance

    constructor(client: AxiosInstance) {
        this.client = client
    }

    async all(param: any): Promise<any> {
        return this.client.get('/room/all', param)
    }

    async add(param: any): Promise<any> {
        return this.client.post('/room/add', param)
    }

    async del(param: any): Promise<any> {
        return this.client.post('/room/del', param)
    }
}