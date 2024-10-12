interface WSConfig {
    url: string,
    error?: Function
}

export enum Event {
    Open = 'open',
    Join = 'join',
    Joined = 'joined',
    OtherJoin = 'other-join'
}

interface Message {
    event: Event,
    content?: string
}

export type {
    WSConfig,
    Message
}