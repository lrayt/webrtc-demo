import type { WSConfig } from './data-types';
export declare class Client {
    private readonly client;
    constructor(cfg: WSConfig);
    on(): void;
    send(msg: any): void;
}
