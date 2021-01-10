import { createStore } from 'vuex'
import { Actions } from '../api/api'

const store = createStore({
  state: {
    userJWT: "",
    socket: null,
    username: "",
    loggedIn: false,
  },
  mutations: {
    userJWT: (state, jwt) => {
      localStorage.setItem('userJWT', jwt)
      state.userJWT = jwt
    },
    initializeStore: (state) => {
      let jwt = localStorage.getItem('userJWT');
      let name = localStorage.getItem('username');
      if (jwt) {
        state.userJWT = jwt;
      }
      if (name) {
        state.username = name;
      }
    },
    username: (state, name) => {
      localStorage.setItem('username', name)
      state.username = name
    },
    socket: (state, socket) => {
      state.socket = socket
    },
    loggedIn: (state, loggedIn) => {
      state.loggedIn = loggedIn
    }
  },
  actions: {
    handleMessage: (context, event) => {
      let payload = JSON.parse(event.data)
      switch (payload.Action) {
        case Actions.helloResponse:
          context.commit('userJWT', payload.Args.token)
          context.commit('username', payload.Args.name)
          context.commit('loggedIn', true)
          break
        default:
          console.log("Could not parse message")
          console.log(payload)
          break
      }
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
