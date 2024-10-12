import {ArrayQueue, ConstantBackoff, Websocket, WebsocketBuilder, WebsocketEvent} from 'websocket-ts'
import {Event, Message, WSConfig} from './data-types'

export class Client {
    private readonly client: Websocket
    private readonly onError: Function

    constructor(cfg: WSConfig) {
        this.onError = cfg.error || console.error

        this.client = new WebsocketBuilder(cfg.url)
            .withBuffer(new ArrayQueue())
            .withBackoff(new ConstantBackoff(1000))
            .build()
        this.client.addEventListener(WebsocketEvent.open, () => this.on({event: Event.Open}))
        // this.client.addEventListener(WebsocketEvent.message, (i: Websocket, ev: MessageEvent) => {
        //     console.log('message---->', ev.data);
        // })
        // this.client.addEventListener(WebsocketEvent.error, () => cfg.error())
        // this.client.addEventListener(WebsocketEvent.close, () => {
        //     console.log('close---->');
        // })
    }

    private on(msg: Message) {
        switch (msg.event) {
            case Event.Open: {
                this.send({event: Event.Join})
                break
            }
            default: {
                this.onError(`unknown event: ${msg.event}`)
            }
        }
    }

    private send(msg: Message) {
        this.client.send(JSON.stringify(msg))
    }
}