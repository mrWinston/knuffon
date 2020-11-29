import { createStore } from 'vuex'

const store = createStore({
  state: {
    userJWT: "",
    socket: null
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
    },
    socket: (state, socket) => {
      state.socket = socket
    }
  },
  actions: {
    handleMessage: (context, event) => {
      console.log("Received Message:")
      console.log(event)
    },
  },
  modules: {
  }
});
store.commit('initializeStore')

var socket = new WebSocket('ws://localhost:8000/ws')
socket.onmessage = (event) => {
  store.dispatch('handleMessage', event)
}
socket.onopen = (event) => {
  console.log(event)
  console.log("Successfully connected to the echo websocket server...")
}
store.commit('socket', socket)

export default store
