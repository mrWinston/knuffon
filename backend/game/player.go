package game

import (
	"errors"
	"fmt"
)

type Player struct {
	Name       string
	ID         string
	scoreBoard map[ResultType]int
	bonus      int
	total      int
}

func NewPlayer(name string, id string) *Player {
	scoreBoard := map[ResultType]int{}
	for _, rs := range AllResults {
		scoreBoard[rs] = -1
	}

	return &Player{
		Name:       name,
		ID:         id,
		scoreBoard: scoreBoard,
		bonus:      0,
		total:      0,
	}
}

func (player *Player) AddScore(rt ResultType, score int) error {
	if player.scoreBoard[rt] != -1 {
		return errors.New(fmt.Sprintf("Result '%s' has already been saved.", rt))
	}
	player.scoreBoard[rt] = score
	player.recalculateBonus()
	return nil
}

func (player *Player) getTotalScore() int {
	total := 0
	for _, score := range player.scoreBoard {
		if score != -1 {
			total += score
		}
	}
	return total + player.bonus
}

func (player *Player) recalculateBonus() {
	total := 0

	for rt, score := range player.scoreBoard {
		switch rt {
		case Ones, Twos, Threes, Fours, Fives, Sixes:
			if score != -1 {
				total += score
			}
		}
	}
	if total >= BONUS_THRESHOLD {
		player.bonus = BONUS_SCORE
	}
}

func (player *Player) FilterRollResult(rs *RollResult) {
	for playerRoll, score := range player.scoreBoard {
		if score != -1 {
			delete(rs.Result, playerRoll)
		}
	}
}

func (player *Player) String() string {
	return fmt.Sprintf("Player('%s - %s')", player.ID, player.Name)
}
