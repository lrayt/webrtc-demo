/// <reference types="vite/client" />

interface GlobalConfig {
    serverUrl: string
}

declare interface Window {
    _global_config: GlobalConfig
}