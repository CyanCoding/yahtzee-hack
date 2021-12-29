// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	board := GenerateBoard()
	score := 0

	// Infinite loop until the game ends
	for i := 1; i <= 13; i++ {
		score = CalculateScore(board)
		color.Set(color.FgHiWhite)
		fmt.Print("\033[H\033[2J")
		fmt.Println("CyanCoding's Yahtzee Hack!")

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
					break
				}
			}
			fmt.Print("\033[H\033[2J")
			// They rolled at least once
			if lastRoll[0] != 0 {
				fmt.Printf(
					"Hand: %d %d %d %d %d\n",
					lastRoll[0],
					lastRoll[1],
					lastRoll[2],
					lastRoll[3],
					lastRoll[4],
				)
				// Ask the user which option they want to fill out on their board
				fmt.Println("Please pick one of the following to fill in.")
				options := FindPossibleOptions(board, lastRoll)

				// Until they pick a valid option
				for true {
					fmt.Print("Option number > ")
					_, _ = fmt.Scanln(&input)

					// If it's not a valid input number
					if (input - 1) >= len(options) {
						color.Set(color.FgHiRed)
						fmt.Println("Invalid option!")
						color.Set(color.FgHiWhite)
					} else { // It is a valid input number
						if (input - 1) == (len(options) - 1) { // Cross out selected
							board = CrossOut(board)
						} else { // Regular option selected
							for j := 0; j < (len(board) - 1); j++ {
								if board[j].id == options[input-1].id {
									board[j].points = options[input-1].points
								}
							}
						}
						break // Only break when they've input a valid input number
					}
				}
			} else {
				// We do this because they didn't even roll once
				i--
			}
		} else if input == 3 { // Calculate score

		} else {
			fmt.Println("Invalid action. Try again.")
			i-- // We don't want to waste a turn on a bad action
			continue
		}
	}
}
