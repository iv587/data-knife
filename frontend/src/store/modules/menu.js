const state = () => {
  return {
    activeIndex: '/redis/connections'
  }
}

const mutations = {
  setActiveIndex(state, activeIndex) {
    state.activeIndex = activeIndex
  }
}

const actions = {
  setMenuActiveIndex({ commit }, activeIndex) {
    commit('setActiveIndex', activeIndex)
  }
}

export default {
  state,
  actions,
  mutations
}
