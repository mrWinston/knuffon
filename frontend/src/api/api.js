const actions = {
  hello: "hello",
  joinGame: "join_game",
  hostGame: "host_game",
  startGame: "start_game"
}

export default Api = {
  loginJson: (username, token) => {
    return {
      Action: actions.hello,
      Token: token,
      Args: {
        name: username
      }
    }
  } 
}
