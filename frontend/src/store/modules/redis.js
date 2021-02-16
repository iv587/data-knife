import { connectionInfo } from '@/api/redis'

const state = () => {
  return {
    current: {
      id: 0,
      dbNo: 0,
      name: ''
    }
  }
}

const mutations = {
  setCurrent(state, current) {
    state.current = current
  },
  setDbNo (state, dbNo) {
    state.current.dbNo = dbNo
  },
  resetCurrent(state, id) {
    if (state.current.id === id) {
      state.current.id = 0
      state.current.dbNo = 0
      state.current.dbNo = ''
    }
  }
}

const actions = {
  currentServer({ commit }, current) {
    commit('setCurrent', current)
  },
  currentServer1({ commit }, id) {
    commit('resetCurrent', id)
  },
  setDbNo({ commit } , dbNo) {
    commit('setDbNo', dbNo)
  }

}
export default {
  state,
  actions,
  mutations
}
