package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
	"github.com/fatih/color"
)

func GenerateBoard() *orderedmap.OrderedMap {
	board := orderedmap.NewOrderedMap()

	board.Set("1's", 0)
	board.Set("2's", 0)
	board.Set("3's", 0)
	board.Set("4's", 0)
	board.Set("5's", 0)
	board.Set("6's", 0)
	board.Set("Three-of-a-kind", 0)
	board.Set("Four-of-a-kind", 0)
	board.Set("Full house", 0)
	board.Set("Small straight", 0)
	board.Set("Large straight", 0)
	board.Set("Yahtzee", 0)
	board.Set("Chance", 0)

	return board
}

func FindPossibleOptions(board *orderedmap.OrderedMap, dice [5]int) {
	points := 0
	for i := 0; i < 5; i++ {
		points += dice[i]
	}

	m := DiceMap(dice)
	possibilities := orderedmap.NewOrderedMap()

	// Calculate upper hand
	value, _ := board.Get("1's")
	if m[1] > 0 && value == 0 {
		possibilities.Set("1's", m[1])
	}
	value, _ = board.Get("2's")
	if m[2] > 0 && value == 0 {
		possibilities.Set("2's", m[2]*2)
	}
	value, _ = board.Get("3's")
	if m[3] > 0 && value == 0 {
		possibilities.Set("3's", m[3]*3)
	}
	value, _ = board.Get("4's")
	if m[4] > 0 && value == 0 {
		possibilities.Set("4's", m[4]*4)
	}
	value, _ = board.Get("5's")
	if m[5] > 0 && value == 0 {
		possibilities.Set("5's", m[5]*5)
	}
	value, _ = board.Get("6's")
	if m[6] > 0 && value == 0 {
		possibilities.Set("6's", m[6]*6)
	}

	// Calculate lower hand
	value, _ = board.Get("Three-of-a-kind")
	if CalculateThreeKind(m) == 1 && value == 0 {
		possibilities.Set("Three-of-a-kind", points)
	}
	value, _ = board.Get("Four-of-a-kind")
	if CalculateFourKind(m) == 1 && value == 0 {
		possibilities.Set("Four-of-a-kind", points)
	}
	value, _ = board.Get("Full house")
	if CalculateFullHouse(m) == 1 && value == 0 {
		possibilities.Set("Full house", 25)
	}
	value, _ = board.Get("Small straight")
	if CalculateSmallStraight(dice) == 1 && value == 0 {
		possibilities.Set("Small straight", 30)
	}
	value, _ = board.Get("Large straight")
	if CalculateLargeStraight(dice) == 1 && value == 0 {
		possibilities.Set("Large straight", 40)
	}
	if CalculateYahtzee(m) == 1 {
		// A Yahtzee is worth +50 points each time.
		// The first is worth 50, then 100, then 150, etc.
		value, _ = board.Get("Yahtzee")
		yahtzeeValue := value.(int) + 50
		possibilities.Set("Yahtzee", yahtzeeValue)
	}

	color.Set(color.FgHiYellow)
	i := 1
	for _, key := range possibilities.Keys() {
		val, _ := possibilities.Get(key)
		fmt.Printf("%d. %s (%d", i, key, val)
		if val == 1 {
			fmt.Println(" point)")
		} else {
			fmt.Println(" points)")
		}
		i++
	}

	color.Set(color.FgHiRed)
	fmt.Printf("%d. Cross out\n", i)
	color.Set(color.FgHiWhite)
}
