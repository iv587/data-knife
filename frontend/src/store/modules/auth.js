const state = () => {
  return {
    token: '',
  }
}

const mutations = {
  setToken(state, token) {
    state.token = token
  }
}

const actions = {
  setToken({ commit }, token) {
    return new Promise(resolve => {
      commit('setToken', token)
      resolve()
    })
  },
  removeToken({ commit }) {
    return new Promise(resolve => {
      commit('setToken', '')
      resolve()
    })
  }
}

export default {
  state,
  actions,
  mutations
}
