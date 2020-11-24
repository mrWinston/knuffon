<template>
  <input v-model="gameid" placeholder="GameID">
</template>
<script>
export default {
  name: 'Login',
  data: function() {
    return {
      connection: null
    }
  },
  created: function() {
    this.$store.commit('initializeStore');
    console.log("Initializing WS Connection")
    this.connection = new WebSocket("ws://localhost:8000/ws")
  }
}
</script>
