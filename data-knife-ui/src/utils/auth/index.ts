// @ts-ignore
import Cookies from 'js-cookie'

const TOKEN_KEY = "token"

class TokenUtils {

    getToken() {
        return Cookies.get(TOKEN_KEY)
        // return '15117973172'
    }

    setToken(token: string) {
        Cookies.set(TOKEN_KEY, token, {expires: 7})
    }

    removeToken() {
        Cookies.remove(TOKEN_KEY)
    }

    toLogin () {
        window.location.href = '/login/'
    }
}


const tokenUtils = new TokenUtils()

export default tokenUtils