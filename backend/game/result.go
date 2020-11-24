package game

import (
	"fmt"
	"sort"
)

const (
	FULL_HOUSE_SCORE     int = 25
	SMALL_STRAIGHT_SCORE     = 30
	STRAIGHT_SCORE           = 40
	KNIFFEL_SCORE            = 50
	BONUS_SCORE              = 35
	BONUS_THRESHOLD          = 63
)

type RollResult struct {
	Dice   []int
	Result map[ResultType]int
}

func GetRollResult(dice []int) RollResult {
	res := RollResult{
		Dice:   dice,
		Result: make(map[ResultType]int),
	}

	res.Result[Ones] = countOccurences(dice, 1)
	res.Result[Twos] = countOccurences(dice, 2) * 2
	res.Result[Threes] = countOccurences(dice, 3) * 3
	res.Result[Fours] = countOccurences(dice, 4) * 4
	res.Result[Fives] = countOccurences(dice, 5) * 5
	res.Result[Sixes] = countOccurences(dice, 6) * 6

	res.Result[ThreeOfAKind] = nOfAKindScore(dice, 3)
	res.Result[FourOfAKind] = nOfAKindScore(dice, 4)
	res.Result[FullHouse] = fullHouseScore(dice)

	maxStraight := longestStraight(dice)
	res.Result[SmallStraight] = 0
	res.Result[Straight] = 0
	if maxStraight >= 4 {
		res.Result[SmallStraight] = SMALL_STRAIGHT_SCORE
	}
	if maxStraight >= 5 {
		res.Result[Straight] = STRAIGHT_SCORE
	}

	res.Result[Kniffel] = 0
	if nOfAKindScore(dice, 5) > 0 {
		res.Result[Kniffel] = KNIFFEL_SCORE
	}

	res.Result[Chance] = sumDice(dice)

	return res
}

func countOccurences(dice []int, number int) int {
	count := 0
	for _, die := range dice {
		if die == number {
			count++
		}
	}
	return count
}

func nOfAKindScore(dice []int, n int) int {
	for i := 1; i <= 6; i++ {
		if countOccurences(dice, i) >= n {
			return sumDice(dice)
		}
	}
	return 0
}

func sumDice(dice []int) int {
	sum := 0
	for _, num := range dice {
		sum += num
	}
	return sum
}

func fullHouseScore(dice []int) int {
	two := false
	three := false

	for i := 1; i <= 6; i++ {
		if countOccurences(dice, i) == 3 {
			three = true
		} else if countOccurences(dice, i) == 2 {
			two = true
		}
	}
	if two && three {
		return FULL_HOUSE_SCORE
	}
	return 0
}

func longestStraight(dice []int) int {
	copyied := append([]int{}, dice...)
	fmt.Printf("%+v\n", copyied)
	sort.Ints(copyied)

	curLen := 1
	maxLen := 1
	for i := 1; i < len(copyied); i++ {
		if copyied[i]-copyied[i-1] == 1 {
			curLen++
		} else {
			curLen = 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
