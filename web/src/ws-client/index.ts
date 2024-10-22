import {ArrayQueue, ConstantBackoff, Websocket, WebsocketBuilder, WebsocketEvent} from 'websocket-ts'
import type {Message, Option} from './data-types'
import {Actions} from './data-types'

export class WSClient {
    private readonly client: Websocket
    private readonly onError: Function
    private readonly onClose: Function

    constructor(cfg: Option) {
        this.onError = cfg.error || console.error
        this.onClose = cfg.close || console.log

        this.client = new WebsocketBuilder(cfg.url)
            .withBuffer(new ArrayQueue())
            .withBackoff(new ConstantBackoff(1000))
            .build()
        this.client.addEventListener(WebsocketEvent.open, () => this.on({action: Actions.Open}))
        this.client.addEventListener(WebsocketEvent.message, (i: Websocket, ev: MessageEvent) => this.on(JSON.parse(ev.data)))
        this.client.addEventListener(WebsocketEvent.error, (i, ev) => this.onError(ev))
        this.client.addEventListener(WebsocketEvent.close, () => {
            this.onClose('client is close')
            this.client.close()
        })
    }

    private on(msg: Message) {
        console.log('msg', msg);
        switch (msg.action) {
            case Actions.Open: {
                this.send({action: Actions.Join})
                break
            }
            case Actions.Joined: {
                this.send({action: Actions.Join})
                break
            }
            case Actions.Error: {
                this.onError(msg.content)
                break
            }
            default: {
                this.onError(`unknown event: ${msg.action}`)
            }
        }
    }

    private send(msg: Message) {
        this.client.send(JSON.stringify(msg))
    }
}