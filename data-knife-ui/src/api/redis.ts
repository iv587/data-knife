import httpClient from '@/utils/http'
import {RedisInfoSection, RedisKeyDetail, RedisKeyListRes} from "@/type";
import {baseURL} from "@/api/base";

export async function listDb() {
    const req  = {
        id: 1,
    }
    return httpClient.post(`${baseURL}/api/redis/db/list`, req)
}

export async function redisInfoApi(data?: any)  {
    return httpClient.post<RedisInfoSection[]>(`${baseURL}/api/redis/info`, data)
}

export async function redisDbsApi(data?: any)  {
    return httpClient.post<{list: number[]}>(`${baseURL}/api/redis/dbs`, data)
}

export async function redisKeysApi(data?: any)  {
    return httpClient.post<RedisKeyListRes>(`${baseURL}/api/redis/keys`, data)
}

export async function redisKeyInfoApi(data?: any)  {
    return httpClient.post<RedisKeyDetail>(`${baseURL}/api/redis/key/info`, data)
}

export async function redisKeyUpdateApi(data?: any)  {
    return httpClient.post(`${baseURL}/api/redis/key/update`, data)
}