import {ArrayQueue, ConstantBackoff, Websocket, WebsocketBuilder, WebsocketEvent} from 'websocket-ts'
import type {Message, Option} from './data-types'
import {Actions} from './data-types'

export class WSClient {
    private uid: string = "local"
    private readonly client: Websocket
    private readonly onLog: Function

    constructor(cfg: Option) {
        this.onLog = cfg.log || console.log

        this.client = new WebsocketBuilder(cfg.url)
            .withBuffer(new ArrayQueue())
            .withBackoff(new ConstantBackoff(1000))
            .build()
        this.client.addEventListener(WebsocketEvent.open, () => this.on({action: Actions.Open}))
        this.client.addEventListener(WebsocketEvent.message, (i: Websocket, ev: MessageEvent) => this.on(JSON.parse(ev.data)))
        this.client.addEventListener(WebsocketEvent.error, (i, ev) => this.on({
            action: Actions.Error,
            content: ev.type
        }))
        this.client.addEventListener(WebsocketEvent.close, () => {
            this.on({action: Actions.Close})
        })
    }

    private log(message: Message): void {
        if (!message.uid) {
            message.uid = this.uid
        }
        if (!message.date) {
            message.date = '2023-09-09T00:00:00.000Z'
        }
        this.onLog(message)
    }

    private on(msg: Message) {
        this.log(msg)
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
                // this.onError(msg.content)
                break
            }
            default: {
                // this.onError(`unknown event: ${msg.action}`)
            }
        }
    }

    private send(msg: Message) {
        this.client.send(JSON.stringify(msg))
    }
}