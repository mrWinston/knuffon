<template>
  <div v-if="!loading">
    <form v-on:submit.prevent="joinGame">
      <input v-model="gameID" placeholder="Game ID">
      <input type="submit" value="Join Game">
    </form>
    <form v-on:submit.prevent="hostGame">
      <p>Alternatively:</p>
      <input type="submit" value="Host Game">
    </form>
  </div>
  <div v-else>
    logging in...
  </div>
</template>
<script>
//import Api from '../api/api.js'
export default {
  name: 'GameSelect',
  data: function() {
    return {
      username: this.$store.state.username,
      gameID : "",
      loading: false
    }
  },
  computed: {
    loggedIn () {
      return this.$store.state.loggedIn;
    }
  },
  created: function() {
    console.log("Initilized GameComponent")
  },
  methods: {
    joinGame () {
      console.log("Game Joined")
    },
    hostGame() {
      console.log("Game Hosted")
    }
  },
  watch: {
  }
}
</script>
