import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'
import { getToken } from './auth'
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 20000
})

service.interceptors.request.use(
  config => {
    const fd = new FormData()
    for (const i in config.data) {
      fd.append(i, config.data[i])
    }
    fd.append('token', getToken())
    config.data = fd
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    if (res.code == 0) {
      Message.error({
        message: res.msg,
        duration: 2000
      })
      return Promise.reject(new Error(res.message || 'Error'))
    }
    if (res.code == 3) {
      Message.error({
        message: res.msg,
        duration: 2000,
        onClose: function () {
          store.dispatch('removeToken').then(res => {
            location.reload()
          })
        }
      })
      return Promise.reject(new Error(res.message || 'Error'))
    }
    return res
  },
  error => {
    Message.error({
      message: error,
      duration: 2000
    })
    return Promise.reject(error)
  }
)

export default service
