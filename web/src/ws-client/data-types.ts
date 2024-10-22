interface Option {
    url: string,
    error?: Function
    close?: Function
}

export enum Actions {
    Open = 'open',
    Join = 'join',
    Joined = 'joined',
    OtherJoin = 'other-join',
    Error = 'error'
}

interface Message {
    action: Actions,
    content?: string
}

export type {
    Option,
    Message
}