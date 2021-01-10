package manager

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mrWinston/knuffon/backend/api"
	"github.com/mrWinston/knuffon/backend/game"
	"github.com/mrWinston/knuffon/backend/util"
	log "github.com/sirupsen/logrus"
)

const (
	NUM_WORKER = 2
)

type RawMessage struct {
	Raw  []byte
	Conn *websocket.Conn
}

type lobby struct {
	players []*game.Player
	code    string
}

type GameManager struct {
	Games        map[string]*game.Game
	Players      map[string]*game.Player
	MessageQueue chan *RawMessage
	stopChan     chan bool
	playersMutex *sync.Mutex
	gamesMutex   *sync.Mutex
}

func NewGameManager() *GameManager {
	gm := &GameManager{
		Games:        map[string]*game.Game{},
		Players:      map[string]*game.Player{},
		MessageQueue: make(chan *RawMessage),
		stopChan:     make(chan bool),
		playersMutex: &sync.Mutex{},
		gamesMutex:   &sync.Mutex{},
	}

	log.Infof("Starting Message Handlers with %d Handlers", NUM_WORKER)
	for i := 0; i < NUM_WORKER; i++ {
		go gm.ListenForMessage()
	}
	return gm
}

func (gm *GameManager) searchPlayerByID(id string) *game.Player {
	for _, player := range gm.Players {
		if player.ID == id {
			return player
		}
	}
	return nil
}

func (gm *GameManager) addPlayer(player *game.Player) bool {
	gm.playersMutex.Lock()
	defer gm.playersMutex.Unlock()
	if _, ok := gm.Players[player.ID]; ok {
		return false
	}
	gm.Players[player.ID] = player
	return true
}

func (gm *GameManager) handleMessage(msg api.Message, conn *websocket.Conn) {
  // error handling first
	if msg.Token == "" && msg.Action != api.ACTION_HELLO {
		// TODO: Return a Proper error
		log.Error("Got a Message without token that's not Hello")
		return
	}
  if err:= api.ValidateActionAndParameter(msg.Action, msg.Args); err != nil {
    log.WithFields(GetErrorFields(err)).Error("Received an invalid message")
    return
  }

  if msg.Action == api.ACTION_HELLO {
		playerName := msg.Args["name"]

		userID, err := ValidateToken(msg.Token)
		if err == nil {
			player = gm.Players[userID] // might just return nil
		}
		if err != nil || player == nil { // No Token or No Player
			log.Debug("Got an Hello Message without a Token. Genrating new User")
			userID = util.GeneratePlayerID()
			player = game.NewPlayer(playerName, userID)
			for !gm.addPlayer(player) { // generate New IDs until we find one that fits
				userID = util.GeneratePlayerID()
				player = game.NewPlayer(playerName, userID)
			}
		}
		// from here, we have a valid player
		token, err := CreateToken(player.ID)
		answer := api.Message{
			Action: api.ACTION_HELLO_RESPONSE,
			Args: map[string]string{
				"name":  player.Name,
				"token": token,
			},
		}
		err = conn.WriteJSON(answer)
		if err != nil {
			log.WithFields(GetErrorFields(err)).Error("Could not Marshal Hello Answer")
		}
}

func (gm *GameManager) Stop() {
	gm.stopChan <- true
}

func (gm *GameManager) AddMessage(raw *RawMessage) {
	gm.MessageQueue <- raw
}

func (gm *GameManager) ListenForMessage() {
	for {
		select {
		case <-gm.stopChan:
			log.Error("Stop Waiting for Messages")
			return
		case raw := <-gm.MessageQueue:
			msg, err := api.ParseMessage(raw.Raw)
			if err != nil {
				log.WithFields(GetErrorFields(err)).Error("Error during message parsing")
			} else {
				log.WithFields(GetMsgFields(msg)).Info("Parsed a Message Successfully")
				go gm.handleMessage(msg, raw.Conn)
			}
		}
	}
}

func GetErrorFields(err error) log.Fields {
	return log.Fields{
		"error": err,
	}
}
func GetMsgFields(msg api.Message) log.Fields {
	return log.Fields{
		"msg": msg,
	}
}
