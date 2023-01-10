import axios from 'axios'
import {AxiosRequestConfig} from "axios/index"
import {ApiRes} from "@/utils/http/type";
import auth from "@/utils/auth";
import qs from 'qs'
import message from "@/utils/message";

class ApiError extends Error {
    code: number
    msg: string

    constructor(code: number, msg: string) {
        super(msg);
        this.code = code
        this.msg = msg
    }
}

class ApiHttpClient {

    async post<T>(url: string, data: any, customerErrHandler?: boolean, customerLogin?: boolean) {
        const token = auth.getToken()
        const tokenObj = {
            token: token
        }
        let reqData: any = {
            ...tokenObj,
            ...data
        }
        const options: AxiosRequestConfig = {
            method: 'POST',
            headers: {'content-type': 'application/x-www-form-urlencoded'},
            data: qs.stringify(reqData),
            url,
        }
        return await this.doPost<T>(options, customerErrHandler, customerLogin)
    }

    async doPost<T>(options: AxiosRequestConfig, customerErrHandler?: boolean, customerLogin?: boolean): Promise<ApiRes<T>> {
        try {
            const res = await axios<ApiRes<T>>(options)
            const resData = res.data
            const code = resData.code
            if (code == 1) {
                // 成功
                return Promise.resolve(resData)
            } else {
                if (!customerErrHandler) {
                    throw new ApiError(resData.code, resData.msg ?? '')
                }
                return Promise.resolve(resData)
            }
        } catch (err) {
            let code = 0;
            let msg = '网络错误'
            if (err instanceof ApiError) {
                code = (err as ApiError).code
                msg = (err as ApiError).msg
                if (code == 3 && msg == '') {
                    msg = '登录状态已经失效！'

                } else if (msg == '') {
                    msg = '请求出错了~'
                }
            }
            await message.error(msg)
            if (code == 3) {
                window.location.reload()
            }
            throw err
        }
    }
}

const httpClient = new ApiHttpClient()

export default httpClient

