import {ArrayQueue, ConstantBackoff, Websocket, WebsocketBuilder, WebsocketEvent} from 'websocket-ts'
import type {Message, Option} from './data-types'
import {Actions} from './data-types'
import {Logger} from '@/common/logger'

export class WSClient {
    private uid: string = "local"
    private readonly client: Websocket
    private readonly log: Logger
    private readonly onUserJoin: Function

    constructor(cfg: Option) {
        // init logger
        this.log = new Logger(this.uid, cfg.onLog)
        this.onUserJoin = cfg.onUserJoin || console.log
        // init client
        this.client = new WebsocketBuilder(cfg.url)
            .withBuffer(new ArrayQueue())
            .withBackoff(new ConstantBackoff(1000))
            .build()

        // add Listener
        this.client.addEventListener(WebsocketEvent.open, () => {
            this.log.info('open websocket connect')
            this.on({action: Actions.Open})
        })
        this.client.addEventListener(WebsocketEvent.message, (i: Websocket, ev: MessageEvent) => {
            const msg: Message = JSON.parse(ev.data)
            this.log.message(msg)
            this.on(msg)
        })
        this.client.addEventListener(WebsocketEvent.error, (i, ev) => {
            this.log.error(JSON.stringify(ev))
            this.on({
                action: Actions.Error,
                content: ev.type
            })
        })
        // this.client.addEventListener(WebsocketEvent.close, () => {
        //     this.on({action: Actions.Close})
        // })
    }

    private on(msg: Message) {
        switch (msg.action) {
            case Actions.Open: {
                this.send({action: Actions.Join})
                break
            }
            case Actions.Joined: {
                this.uid = msg.uid || this.uid
                console.log('--->', this.uid);
                break
            }
            case Actions.OtherJoin: {
                this.onUserJoin(msg.uid)
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

    public getUID(): string {
        return this.uid
    }

    public send(msg: Message) {
        this.log.message(msg)
        this.client.send(JSON.stringify(msg))
    }
}