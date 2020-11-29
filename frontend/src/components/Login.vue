<template>
  <input v-model="username" placeholder="Username">
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
    console.log("Initilized Login component");
  }
}
</script>
