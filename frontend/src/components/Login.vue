<template>
  <div v-if="!loginPressed">
    <form v-on:submit.prevent="login">
      <input v-model="username" placeholder="Username">
      <input type="submit" value="OK">
    </form>
  </div>
  <div v-else>
    logging in...
  </div>
</template>
<script>
import Api from '../api/api.js'
export default {
  name: 'Login',
  data: function() {
    return {
      username: this.$store.state.username,
      loginPressed: false
    }
  },
  computed: {
    loggedIn () {
      return this.$store.state.loggedIn
    }
  },
  created: function() {
    console.log("Initilized Login component");
  },
  methods: {
    login() {
      this.loginPressed = true
      let loginJson = Api.loginJson(this.username, this.$store.state.userJWT)
      this.$store.state.socket.send(JSON.stringify(loginJson))
      console.log("Yeah! Logged in!")
    }
  },
  watch: {
    loggedIn(newValue, oldValue) {
      if (newValue) {
        console.log(newValue, oldValue) 
        this.$router.push('/gameSelect')
      } 
    }
  }

}
</script>
