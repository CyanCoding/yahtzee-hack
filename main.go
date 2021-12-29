// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import (
	"fmt"
)

func main() {
	fmt.Println("CyanCoding's Yahtzee Hack!")

	board := GenerateBoard()
	score := 0

	// Infinite loop until the game ends
	for i := 1; i <= 13; i++ {
		fmt.Println("round", i, "- score", score) // round 0 - score 0
		fmt.Println("Actions:")
		fmt.Println("1. View current hand")
		fmt.Println("2. Enter dice")
		fmt.Println("3. Calculate score")
		fmt.Print("Please enter a number > ")

		var input int
		_, _ = fmt.Scanln(&input)

		if input == 1 { // View current hand

		} else if input == 2 { // Enter dice
			var lastRoll, dice [5]int
			fmt.Print("\033[H\033[2J")
			for j := 0; j < 3; j++ { // Up to three rolls per turn
				dice = InputDice()
				if dice[0] != 0 {
					lastRoll = dice
					CalculateLowerHand(dice)
				} else {
					i--
					break
				}
			}
			fmt.Print("\033[H\033[2J")
			if lastRoll[0] != 0 {
				fmt.Printf(
					"Hand: %d %d %d %d %d\n",
					lastRoll[0],
					lastRoll[1],
					lastRoll[2],
					lastRoll[3],
					lastRoll[4],
				)
				fmt.Println("Please pick one of the following to fill in.")
				FindPossibleOptions(board, lastRoll)
				fmt.Print("Option number > ")
				_, _ = fmt.Scanln(&input)
			}
		} else if input == 3 { // Calculate score

		} else {
			fmt.Println("Invalid action. Try again.")
			i-- // We don't want to waste a turn on a bad action
			continue
		}

	}
}
