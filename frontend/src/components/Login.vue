<template>
  <form action="" method="" accept-charset="utf-8">
    <input v-model="username" placeholder="Username">
    <input type="submit" value="OK" v-on:click="login">
  </form>
</template>
<script>
import Api from '../api/api.js'
export default {
  name: 'Login',
  data: function() {
    return {
      username: ""
    }
  },
  created: function() {
    console.log("Initilized Login component");
  },
  methods: {
    login: (event) =>{
      this.$state.state.socket.send(JSON.stringify({
        
      }))
      console.log("Yeah! Logged in!")
      console.log(event)
    }
  }

}
</script>
