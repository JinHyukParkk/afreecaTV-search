// state
const state = {
    message: 'Hello'
}

// mutations
const mutations = {
     changeMessage (state, newMsg) {
      state.message = newMsg
    }
}

// actions
const actions = {
    callMutation ({ state, commit }, { newMsg }) {
      commit('changeMessage', newMsg)
    }
}

// getters
const getters = {
    getMsg (state) {
      return `${state.message} => Length : ${state.message.length}`
    }
}

export default {
  state,
  mutations,
  actions,
  getters
}