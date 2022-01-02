package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// fillInOtherDice randomly rolls the dice you don't want to keep.
// keep[] has the numbers you do want to keep.
func fillInOtherDice(dice [5]int, keep []int) [5]int {
	for i := 0; i < len(dice); i++ {
		erase := true
		for j := 0; j < len(keep); j++ {
			if dice[i] == keep[j] {
				erase = false
			}
		}
		// There were no matches in dice[i] with keep
		if erase {
			rand.Seed(time.Now().UnixNano())
			dice[i] = rand.Intn(7) // 0 - 6
		}
	}
	return dice
}

// This tells us what numbers we need in order to get a large straight
func calculateRemainderLargeStraight(dice [5]int) (keepDice [4]int) {
	possibility1 := "12345"
	possibility2 := "23456"

	keepDiceInt := 0
	for i := 0; i < 5; i++ {
		oneExecuted := false // Used to determine if either option occurs
		character := strconv.Itoa(dice[i])
		if strings.Contains(possibility1, character) {
			possibility1 = strings.Replace(possibility1, character, "", -1)
			oneExecuted = true
		}
		if strings.Contains(possibility2, character) {
			possibility2 = strings.Replace(possibility2, character, "", -1)
			oneExecuted = true
		}

		if oneExecuted {
			// They already have a full straight because it's trying to add more than
			// 4 to keep.
			if keepDiceInt == 4 {
				return
			}
			keepDice[keepDiceInt] = dice[i]
			keepDiceInt++
		}
	}
	fmt.Println(keepDice)

	return
}

// InterpretFinish returns the new board and new dice
func InterpretFinish(board [13]ScoreItem, line string, dice [5]int) ([13]ScoreItem, [5]int) {
	// Example: "1. Cross out your Yahtzee."
	line = line[3:]

	m := DiceMap(dice)
	largestKey := 0
	largestValue := 0
	for key, value := range m {
		if value > largestValue {
			largestKey = key
			largestValue = value
		}
	}

	//crossOut := false
	//crossOutOption := ""
	fillInOption := ""
	keepRolling := false
	var keepDice [4]int

	if keepRolling && fillInOption == "" {
		// I did this so I can keep the variables lol
	}

	if line == "Take the Yahtzee and stop rolling." {
		fillInOption = "yahtzee"
	} else if line == "Take your 1's and get the bonus 35 points." {
		fillInOption = "1's"
	} else if line == "Take your 2's and get the bonus 35 points." {
		fillInOption = "2's"
	} else if line == "Take your 3's and get the bonus 35 points." {
		fillInOption = "3's"
	} else if line == "Take your 4's and get the bonus 35 points." {
		fillInOption = "4's"
	} else if line == "Take your 5's and get the bonus 35 points." {
		fillInOption = "5's"
	} else if line == "Take your 6's and get the bonus 35 points." {
		fillInOption = "6's"
	} else if line == "Take the large straight and stop rolling." {
		fillInOption = "straight"
	} else if line == "Take the Full house and stop rolling." {
		fillInOption = "full house"
	} else if line == "Go for a four-of-a-kind." {
		keepDice[0] = largestKey
		keepRolling = true
	} else if line == "Go for a Yahtzee." {
		keepDice[0] = largestKey
		keepRolling = true
	} else if line == "Go for a large straight." {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	} else if line == "" {
		fillInOption = ""
	}

	return board, dice
}
