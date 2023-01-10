import {defineStore} from "pinia";
import {LoginRes, MenuItem, User} from "@/type";
import tokenUtils from "@/utils/auth";

const useUserStore = defineStore('userStore', {
    state: () => {
        return {
            token: '',
            userName: ''
        }
    },
    actions: {
        setUserInfo(info: User) {
            this.userName = info.userName
        },
        setLoginRes(res: LoginRes) {
            this.token = res.token
            tokenUtils.setToken(res.token)
            this.setUserInfo(res)
        },
        logout() {
            this.token = ''
            tokenUtils.removeToken()
        }
    }
})

export default useUserStore