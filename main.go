// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import (
	"fmt"
	"github.com/golang-demos/chalk"
)

var YahtzeeValue int = 0
var YahtzeeMultiplied int = 50

func main() {
	board := GenerateBoard()
	score := 0
	clearScreen := true

	// Infinite loop until the game ends
	for i := 1; i <= 13; i++ {
		score = CalculateTotalScore(board)
		fmt.Print(chalk.Reset())

		if clearScreen {
			fmt.Print("\033[H\033[2J")
		} else {
			clearScreen = true
		}

		fmt.Println("CyanCoding's Yahtzee Hack!")

		fmt.Println("round", i, "- score", score) // round 0 - score 0
		fmt.Println("Actions:")
		fmt.Println("1. View your score card")
		fmt.Println("2. Enter dice")
		fmt.Print("Please enter a number > ")

		var input int
		_, _ = fmt.Scanln(&input)

		if input == 1 { // View current scorecard
			DisplayScoreBoard(board)
			clearScreen = false
			fmt.Println()
			i--
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
				board = FindPossibleOptions(board, lastRoll)
			} else {
				// We do this because they didn't even roll once
				i--
			}
		} else { // Invalid action
			fmt.Println(chalk.RedLight())
			fmt.Print("Invalid action. Try again.")
			fmt.Println(chalk.Reset())
			fmt.Println()
			clearScreen = false
			i-- // We don't want to waste a turn on a bad action
			continue
		}
	}
}
