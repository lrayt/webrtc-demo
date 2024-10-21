import type {AxiosInstance, AxiosResponse} from 'axios'

export default class RoomAPI {
    private readonly client: AxiosInstance;

    constructor(client: AxiosInstance) {
        this.client = client
    }

    loadAll(): Promise<AxiosResponse> {
        return this.client.get('room/all')
    }

    add(param: any): Promise<AxiosResponse> {
        return this.client.post('room/add', param)
    }

    del(param: any): Promise<AxiosResponse> {
        return this.client.post('room/del', param)
    }
}

