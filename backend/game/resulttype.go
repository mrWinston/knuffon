package game

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ResultType string

const (
	Ones          ResultType = "ones"
	Twos                     = "twos"
	Threes                   = "threes"
	Fours                    = "fours"
	Fives                    = "fives"
	Sixes                    = "sixes"
	ThreeOfAKind             = "threeofakind"
	FourOfAKind              = "fourofakind"
	FullHouse                = "fullhouse"
	SmallStraight            = "smallstraight"
	Straight                 = "straight"
	Kniffel                  = "kniffel"
	Chance                   = "chance"
)

var AllResults []ResultType = []ResultType{
	Ones,
	Twos,
	Threes,
	Fours,
	Fives,
	Sixes,
	ThreeOfAKind,
	FourOfAKind,
	FullHouse,
	SmallStraight,
	Straight,
	Kniffel,
	Chance,
}

func (rt *ResultType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	inputResultType := ResultType(s)

	if err := inputResultType.IsValid(); err != nil {
		return err
	}
	*rt = inputResultType
	return nil
}

func (rt *ResultType) IsValid() error {
	if findInSlice(rt, AllResults) >= 0 {
		return nil
	}

	return errors.New(fmt.Sprintf("%v is an Invalid Result Type", rt))
}

func findInSlice(search *ResultType, slice []ResultType) int {
	for i, rs := range slice {
		if rs == *search {
			return i
		}
	}
	return -1
}
