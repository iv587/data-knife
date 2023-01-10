import exp from "constants";

export interface MenuItem {
    id?: string
    title?: string
    icon?: string
    index: string
    children?: MenuItem[]
}

export interface Connection {
    id: number
    name: string
    addr: string
    password: string
}

interface RedisInfoItem {
    name: string
    value: string
}

export interface RedisInfoSection{
    label: string
    items: RedisInfoItem[]
}

export interface RedisKey {
    key: string
    type: string
    ttl: number
}

export interface RedisKeyListRes {
    list: RedisKey[],
    max: number,
    total: number
}

export interface RedisKeyDetail extends RedisKey{
    value: string
    values: RedisKvPair[]
}

export interface RedisKvPair {
    field: string
    val:   string
    score: string
}

export interface User {
    userName: string
}

export interface LoginRes extends User{
    token: string
}