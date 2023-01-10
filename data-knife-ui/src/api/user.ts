import {baseURL} from "@/api/base";
import httpClient from "@/utils/http";
import {LoginRes} from "@/type";

export async function loginApi(data: {userName: string, password: string}) {
    return httpClient.post<LoginRes>(`${baseURL}/api/login`, data)
}

export async function updatePasswordApi(data: {oldPassword: string, newPassword: string}) {
    return httpClient.post(`${baseURL}/api/user/update/password`, updatePasswordApi)
}

export async function appCheckApi(data: any) {
    return httpClient.post<{
        init: number,
    }>(`${baseURL}/api/app/check`, data)
}

export async function appInitApi(data: any) {
    return httpClient.post(`${baseURL}/api/app/init`, data)
}