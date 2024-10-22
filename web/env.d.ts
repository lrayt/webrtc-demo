/// <reference types="vite/client" />
interface GlobalConfig {
    serverUrl: string
    roomUrl: string
}


declare interface Window {
    _globalConfig: GlobalConfig
}
