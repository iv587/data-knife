import Vue from 'vue'
import Vuex from 'vuex'
import redis from './modules/redis'
import menu from '@/store/modules/menu'
import auth from '@/store/modules/auth'
import createPersistedState from 'vuex-persistedstate'
import Cookies from 'js-cookie'

Vue.use(Vuex)

export default new Vuex.Store({
  plugins: [createPersistedState({
    storage: {
      getItem: key => Cookies.get(key),
      setItem: (key, value) => {
        Cookies.set(key, value)
      },
      removeItem: key => Cookies.remove(key)
    }
  })],
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    redis,
    menu,
    auth
  }
})
