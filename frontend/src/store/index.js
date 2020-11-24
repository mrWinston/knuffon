import { createStore } from 'vuex'

export default createStore({
  state: {
    userJWT: ""
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
})
