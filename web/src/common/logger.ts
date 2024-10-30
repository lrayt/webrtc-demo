import {NowDate} from "@/common/date";
import {Actions, type Message} from "@/ws-client/data-types";

export enum LogLevel {
    Info = 0,
    Warning = 1,
    Error = 2
}

interface LogInfo {
    uid?: string
    level: LogLevel
    date?: string
    content: string
}

export type {
    LogInfo
}

const SendPrefix: string = 'send ->> '

export class Logger {
    private readonly uid: string
    private readonly output: Function

    constructor(uid: string, output?: Function) {
        this.uid = uid
        this.output = output || console.log
    }


    info(msg: string, uid?: string) {
        this.output({
            uid: uid || this.uid,
            level: LogLevel.Info,
            date: NowDate(),
            content: msg
        })
    }

    warn(msg: string, uid?: string) {
        this.output({
            uid: uid || this.uid,
            level: LogLevel.Warning,
            date: NowDate(),
            content: msg
        })
    }

    error(msg: string, uid?: string) {
        this.output({
            uid: uid || this.uid,
            level: LogLevel.Error,
            date: NowDate(),
            content: msg
        })
    }

    message(msg: Message) {
        switch (msg.action) {
            case Actions.Join: {
                this.output({
                    uid: msg.uid || this.uid,
                    level: LogLevel.Info,
                    date: NowDate(),
                    content: `${SendPrefix}join room`
                })
                break
            }
            case Actions.Joined: {
                this.output({
                    uid: msg.uid || this.uid,
                    level: LogLevel.Info,
                    date: NowDate(),
                    content: 'you joined room'
                })
                break
            }
            case Actions.OtherJoin: {
                this.output({
                    uid: msg.uid || this.uid,
                    level: LogLevel.Info,
                    date: NowDate(),
                    content: `other user joined`
                })
                break
            }
            case Actions.Error: {
                this.output({
                    uid: msg.uid || this.uid,
                    level: LogLevel.Error,
                    date: NowDate(),
                    content: msg.content
                })
                break
            }
            default: {
                console.log('--->', msg);
            }
        }
    }
}