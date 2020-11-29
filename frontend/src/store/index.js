import { createStore } from 'vuex'

const store = createStore({
  state: {
    userJWT: ""
    socket: null,
  },
  mutations: {
    userJWT: (state, jwt) => {
      localStorage.setItem('userJWT', jwt)
      state.userJWT = jwt
    },
    initializeStore: (state) => {
      let jwt = localStorage.getItem('userJWT');
      if (jwt) {
        state.userJWT = jwt;
      }
    }
  },
  actions: {
  },
  modules: {
  }
});
store.commit('initializeStore')
var socket = new WebSocket('ws://localhost:8000/ws')
store.commit('socket', socket)

export default store
