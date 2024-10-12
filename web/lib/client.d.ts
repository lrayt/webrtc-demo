import { WSConfig } from './data-types';
export declare class Client {
    private readonly client;
    private readonly onError;
    constructor(cfg: WSConfig);
    private on;
    private send;
}
