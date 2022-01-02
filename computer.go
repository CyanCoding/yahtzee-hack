package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// FillInOtherDice randomly rolls the dice you don't want to keep.
// keep[] has the numbers you do want to keep.
func FillInOtherDice(dice [5]int, keep [4]int) [5]int {
	// We can do this because we know we already have everything in [4]keep
	dice[0] = keep[0]
	dice[1] = keep[1]
	dice[2] = keep[2]
	dice[3] = keep[3]

	for i := 0; i < len(dice); i++ {
		if dice[i] == 0 {
			rand.Seed(time.Now().UnixNano())
			dice[i] = rand.Intn(6) + 1 // 1 - 6
		}
	}
	return dice
}

// CalculateRemainderLargeStraight tells us what numbers we need in order to get a large straight
// It returns keepDice, which is an array of numbers we don't need to re-roll.
func CalculateRemainderLargeStraight(dice [5]int) (keepDice [4]int) {
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

	return
}

// calculateRemainderSmallStraight tells us what numbers we need in order to get a small straight.
// It returns the dice we don't need to re-roll.
func calculateRemainderSmallStraight(dice [5]int) (keepDice [4]int) {
	original1 := "1234"
	original2 := "2345"
	original3 := "3456"
	possibility1 := "1234"
	possibility2 := "2345"
	possibility3 := "3456"

	for i := 0; i < 5; i++ {
		character := strconv.Itoa(dice[i])
		if strings.Contains(possibility1, character) {
			possibility1 = strings.Replace(possibility1, character, "", -1)
		}
		if strings.Contains(possibility2, character) {
			possibility2 = strings.Replace(possibility2, character, "", -1)
		}
		if strings.Contains(possibility3, character) {
			possibility3 = strings.Replace(possibility3, character, "", -1)
		}
	}

	usingString := possibility1
	original := original1

	if len(possibility2) < len(possibility1) {
		usingString = possibility2
		original = original2
	} else if len(possibility3) < len(possibility1) {
		usingString = possibility3
		original = original3
	}

	for i := 0; i < len(usingString); i++ {
		if strings.Contains(original, string(usingString[i])) {
			original = strings.Replace(original, string(usingString[i]), "", -1)
		}
	}

	for i := 0; i < len(original); i++ {
		keepDice[i] = int(original[i] - '0')
	}

	return
}

// calculateRemainderFullHouse tells us what numbers we need in order to get a full house.
// It returns the dice we don't need to re-roll
func calculateRemainderFullHouse(m map[int]int) (keepDice [4]int) {
	haveThree := 0
	haveTwo := 0
	keepDiceInt := 0

	for key, value := range m {
		if value == 2 {
			keepDice[keepDiceInt] = key
			keepDiceInt++
			haveTwo++
		} else if value == 3 {
			haveThree++
			keepDice[keepDiceInt] = key
			keepDiceInt++
		}
	}

	return
}

// InterpretFinish returns the new board and new dice
func InterpretFinish(board [13]ScoreItem, line string, dice [5]int) (crossOut bool,
	crossOutOption string,
	fillInOption string,
	keepRolling bool,
	keepDice [4]int) {
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

	// This is for normal results - not crossing out
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
		fillInOption = "large straight"
	} else if line == "Take the Full house and stop rolling." {
		fillInOption = "full house"
	} else if line == "Go for a four-of-a-kind." {
		keepDice[0] = largestKey
		keepRolling = true
	} else if line == "Go for a Yahtzee." {
		keepDice[0] = largestKey
		keepRolling = true
	} else if line == "Go for a large straight." {
		keepDice = CalculateRemainderLargeStraight(dice)
		keepRolling = true
	} else if line == "Go for a three-of-a-kind." {
		keepDice[0] = largestKey
		keepRolling = true
	} else if line == "Take the small straight." {
		fillInOption = "small straight"
	} else if line == "Take your 1's." || line == "Take your 1's and get the bonus 35 points." {
		fillInOption = "1's"
	} else if line == "Take your 2's." || line == "Take your 2's and get the bonus 35 points." {
		fillInOption = "2's"
	} else if line == "Take your 3's." || line == "Take your 3's and get the bonus 35 points." {
		fillInOption = "3's"
	} else if line == "Take your 4's." || line == "Take your 4's and get the bonus 35 points." {
		fillInOption = "4's"
	} else if line == "Take your 5's." || line == "Take your 5's and get the bonus 35 points." {
		fillInOption = "5's"
	} else if line == "Take your 6's." || line == "Take your 6's and get the bonus 35 points." {
		fillInOption = "6's"
	} else if line == "Take the four-of-a-kind." {
		fillInOption = "four-of-a-kind"
	} else if line == "Take the three-of-a-kind." {
		fillInOption = "three-of-a-kind"
	} else if line == "Go for a full house." {
		keepDice = calculateRemainderFullHouse(m)
		keepRolling = true
	} else if line == "Keep your 6's and keep rolling." || line == "Roll for 6's" {
		keepDice[0] = 6
		keepRolling = true
	} else if line == "Keep your 5's and keep rolling." || line == "Roll for 5's" {
		keepDice[0] = 5
		keepRolling = true
	} else if line == "Keep your 4's and keep rolling." || line == "Roll for 4's" {
		keepDice[0] = 4
		keepRolling = true
	} else if line == "Keep your 3's and keep rolling." || line == "Roll for 3's" {
		keepDice[0] = 3
		keepRolling = true
	} else if line == "Keep your 2's and keep rolling." || line == "Roll for 2's" {
		keepDice[0] = 2
		keepRolling = true
	} else if line == "Keep your 1's and keep rolling." || line == "Roll for 1's" {
		keepDice[0] = 1
		keepRolling = true
	} else if line == "Go for a small straight." {
		keepDice = calculateRemainderSmallStraight(dice)
		keepRolling = true
	} else if line == "Re-roll low numbers to get a good chance." {
		keepDice[0] = 4
		keepDice[1] = 5
		keepDice[2] = 6
		keepRolling = true
	} else if line == "Use your chance." {
		fillInOption = "chance"
	}

	// This dictates if we need to cross something out
	if line == "Cross out your Yahtzee." {
		crossOut = true
		crossOutOption = "yahtzee"
	} else if line == "Cross out your four-of-a-kind." {
		crossOut = true
		crossOutOption = "four-of-a-kind"
	} else if line == "Cross out your three-of-a-kind." {
		crossOut = true
		crossOutOption = "three-of-a-kind"
	} else if line == "Cross out your 1's." {
		crossOut = true
		crossOutOption = "1's"
	} else if line == "Cross out your 2's." {
		crossOut = true
		crossOutOption = "2's"
	} else if line == "Cross out your 3's." {
		crossOut = true
		crossOutOption = "3's"
	} else if line == "Cross out your 4's." {
		crossOut = true
		crossOutOption = "4's"
	} else if line == "Cross out your 5's." {
		crossOut = true
		crossOutOption = "5's"
	} else if line == "Cross out your 6's." {
		crossOut = true
		crossOutOption = "6's"
	} else if line == "Cross out your full house." {
		crossOut = true
		crossOutOption = "full house"
	} else if line == "Cross out your large straight." {
		crossOut = true
		crossOutOption = "large straight"
	} else if line == "Cross out your small straight." {
		crossOut = true
		crossOutOption = "small straight"
	}

	return
}
