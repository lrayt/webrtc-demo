interface Option {
    url: string,
    onLog?: Function
    onMessage?: Function
    onUserJoin?: Function
}

export enum Actions {
    Open = 'open',
    Join = 'join',
    Joined = 'joined',
    OtherJoin = 'other-join',
    Message = 'message',
    Error = 'error',
    Leave = 'leave'
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