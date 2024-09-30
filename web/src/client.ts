import {ArrayQueue, ConstantBackoff, Websocket, WebsocketBuilder, WebsocketEvent} from 'websocket-ts'
import type {WSConfig} from './data-types'

export class Client {
    private readonly client: Websocket


    constructor(cfg: WSConfig) {
        this.client = new WebsocketBuilder(cfg.url)
            .withBuffer(new ArrayQueue())
            .withBackoff(new ConstantBackoff(1000))
            .build()
        this.client.addEventListener(WebsocketEvent.open, () => {
            console.log('open---->');
        })
        this.client.addEventListener(WebsocketEvent.message, (i: Websocket, ev: MessageEvent) => {
            console.log('message---->', ev.data);
        })
        this.client.addEventListener(WebsocketEvent.error, () => {
            console.log('error---->');
        })
        this.client.addEventListener(WebsocketEvent.close, () => {
            console.log('close---->');
        })
    }

    on() {

    }

    send(msg: any) {
        this.client.send(JSON.stringify(msg))
    }
}