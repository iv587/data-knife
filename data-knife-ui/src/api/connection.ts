import httpClient from '@/utils/http'
import {Connection} from "@/type";
import {baseURL} from "@/api/base";


export async function listConnection(data?: any)  {
    return httpClient.post<{list: Connection[]}>(`${baseURL}/api/connection/list`, data)
}

export async function testConnection(data?: any) {
    return httpClient.post(`${baseURL}/api/connection/test`, data, true)
}

export async function updateConnection(data?: any) {
    return httpClient.post(`${baseURL}/api/connection/update`, data)
}

export async function getConnection<T>(data: any) {
    return httpClient.post<T>(`${baseURL}/api/connection/info`, data, false)
}