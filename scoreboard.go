package main

import (
	"fmt"
	"github.com/fatih/color"
)

type ScoreItem struct {
	name   string
	points int
}

func GenerateBoard() [13]ScoreItem {
	board := [13]ScoreItem{}

	board[0] = ScoreItem{"1's", 0}
	board[1] = ScoreItem{"2's", 0}
	board[2] = ScoreItem{"3's", 0}
	board[3] = ScoreItem{"4's", 0}
	board[4] = ScoreItem{"5's", 0}
	board[5] = ScoreItem{"6's", 0}
	board[6] = ScoreItem{"Three-of-a-kind", 0}
	board[7] = ScoreItem{"Four-of-a-kind", 0}
	board[8] = ScoreItem{"Full house", 0}
	board[9] = ScoreItem{"Small straight", 0}
	board[10] = ScoreItem{"Large straight", 0}
	board[11] = ScoreItem{"Yahtzee", 0}
	board[12] = ScoreItem{"Chance", 0}

	return board
}

// FindPossibleOptions Finds the playable options from the given dice and scoreboard and prints them out.
// Returns the vector of playable options with "Cross out" in the last position.
func FindPossibleOptions(board [13]ScoreItem, dice [5]int) []ScoreItem {
	points := 0
	for i := 0; i < 5; i++ {
		points += dice[i]
	}

	m := DiceMap(dice)
	//possibilities := orderedmap.NewOrderedMap()
	possibilities := make([]ScoreItem, 0)

	// Calculate upper hand
	// Note: it's m[1] instead of m[0] because 1 references the number, not the position
	if m[1] > 0 && board[0].points == 0 {
		possibilities = append(possibilities, ScoreItem{"1's", m[1]})
	}

	if m[2] > 0 && board[1].points == 0 {
		possibilities = append(possibilities, ScoreItem{"2's", m[2] * 2})
	}

	if m[3] > 0 && board[2].points == 0 {
		possibilities = append(possibilities, ScoreItem{"3's", m[3] * 3})
	}

	if m[4] > 0 && board[3].points == 0 {
		possibilities = append(possibilities, ScoreItem{"4's", m[4] * 4})
	}

	if m[5] > 0 && board[4].points == 0 {
		possibilities = append(possibilities, ScoreItem{"5's", m[5] * 5})
	}

	if m[6] > 0 && board[5].points == 0 {
		possibilities = append(possibilities, ScoreItem{"6's", m[6] * 6})
	}

	// Calculate lower hand
	if CalculateThreeKind(m) == 1 && board[6].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Three-of-a-kind", points})
	}

	if CalculateFourKind(m) == 1 && board[7].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Four-of-a-kind", points})
	}

	if CalculateFullHouse(m) == 1 && board[8].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Full house", 25})
	}

	if CalculateSmallStraight(dice) == 1 && board[9].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Small straight", 30})
	}

	if CalculateLargeStraight(dice) == 1 && board[10].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Large straight", 40})
	}
	if CalculateYahtzee(m) == 1 {
		// A Yahtzee is worth +50 points each time.
		// The first is worth 50, then 100, then 150, etc.
		yahtzeeValue := board[11].points + 50
		possibilities = append(possibilities, ScoreItem{"Yahtzee", yahtzeeValue})
	}
	if board[12].points == 0 { // This is the "Chance" option
		possibilities = append(possibilities, ScoreItem{"Chance", points, 12})
	}

	color.Set(color.FgHiYellow)
	var i int
	for i = 0; i < len(possibilities); i++ {
		fmt.Printf("%d. %s (%d", i+1, possibilities[i].name, possibilities[i].points)
		if possibilities[i].points == 1 {
			fmt.Println(" point)")
		} else {
			fmt.Println(" points)")
		}
	}

	color.Set(color.FgHiRed)
	fmt.Printf("%d. Cross out\n", i+1)
	color.Set(color.FgHiWhite)

	possibilities = append(possibilities, ScoreItem{"Cross out", 0})
	return possibilities
}
