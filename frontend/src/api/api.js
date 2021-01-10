export const Actions = {
  hello: "hello",
  helloResponse: "hello_response",
  joinGame: "join_game",
  hostGame: "host_game",
  startGame: "start_game"

}

var Api = {
  loginJson: (username, token) => {
    return {
      Action: Actions.hello,
      Token: token,
      Args: {
        name: username
      }
    }
  }
}

export default Api
