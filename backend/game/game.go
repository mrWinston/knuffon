package game

import (
	"errors"
	"log"

	"github.com/mrWinston/knuffon/backend/util"
)

type Game struct {
	Players      []*Player
	curPlayerNum int
	rollsDone    int
	dice         []int
	Done         bool
	MaxTurns     int
	turn         int
}

func CreateGame(players []*Player) *Game {
	return &Game{
		Players:      players,
		dice:         []int{0, 0, 0, 0, 0},
		curPlayerNum: 0,
		rollsDone:    0,
		Done:         false,
		MaxTurns:     len(AllResults),
		turn:         1,
	}
}

func (game *Game) HasStarted() bool {
	return game.turn == 0 && game.rollsDone == 0 && game.curPlayerNum == 0
}

func (game *Game) GetCurrentPlayer() *Player {
	return game.Players[game.curPlayerNum]
}

func (game *Game) TurnDone(rt ResultType) error {
	if game.Done {
		return errors.New("Game is already finished")
	}

	score := GetRollResult(game.dice).Result[rt]

	if err := game.GetCurrentPlayer().AddScore(rt, score); err != nil {
		return err
	}

	game.curPlayerNum = (game.curPlayerNum + 1) % len(game.Players)
	log.Printf("CurPlayerNum: %d", game.curPlayerNum)
	log.Printf("CurPlayer: %d", game.GetCurrentPlayer())
	game.rollsDone = 0
	game.dice = []int{0, 0, 0, 0, 0}

	if game.curPlayerNum == 0 {
		game.turn++
	}

	if game.turn > game.MaxTurns {
		game.Done = true
	}

	return nil
}

func (game *Game) Roll(diceSelect []bool) (*RollResult, error) {
	if game.Done {
		return nil, errors.New("Game is already finished")
	}
	if game.rollsDone >= 3 {
		return nil, errors.New("Can not Roll more than 3 Times.")
	}

	for i := 0; i < len(diceSelect); i++ {
		if diceSelect[i] {
			game.dice[i] = util.RollDie()
		}
	}

	game.rollsDone += 1
	rollResult := GetRollResult(game.dice)
	player := game.GetCurrentPlayer()
	player.FilterRollResult(&rollResult)
	return &rollResult, nil

}
