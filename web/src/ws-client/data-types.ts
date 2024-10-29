interface Option {
    url: string,
    log?: Function
    error?: Function
    close?: Function
}

export enum Actions {
    Open = 'open',
    Join = 'join',
    Joined = 'joined',
    OtherJoin = 'other-join',
    Error = 'error',
    Close = 'close'
}

interface Message {
    uid?: string
    date?: string
    action: Actions
    content?: string
}

export type {
    Option,
    Message
}