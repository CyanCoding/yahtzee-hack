// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import (
	"fmt"
	"github.com/golang-demos/chalk"
)

var YahtzeeValue int = 0
var YahtzeeMultiplied int = 50

// EmergencyAdvice This is a variable to use when there's no advice
// but we want to reinforce that the user needs to achieve their goal
var EmergencyAdvice []string

var ActionString string

func main() {
	board := GenerateBoard()
	score := 0
	clearScreen := true

	EmergencyAdvice = make([]string, 0)

	var input int
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

		fmt.Print("round ", i, " - score ", score) // round 0 - score 0
		fmt.Println(chalk.YellowLight())
		fmt.Println("Actions:")
		fmt.Println("1. View your score card")
		fmt.Println("2. Enter dice")
		fmt.Print("3. Run as computer")
		fmt.Println(chalk.Reset())
		fmt.Print("Please enter a number > ")

		if input != 3 {
			_, _ = fmt.Scanln(&input)
		}

		if input == 1 { // View current scorecard
			fmt.Print("\033[H\033[2J")
			DisplayScoreBoard(board)
			clearScreen = false
			fmt.Println()
			i--
		} else if input == 2 || input == 3 { // Enter dice
			var lastRoll, dice [5]int
			fmt.Print("\033[H\033[2J")

			EmergencyAdvice = nil // Clear the slice
			CalculateTargets(board)

			for j := 0; j < 3; j++ { // Up to three rolls per turn
				dice = InputDice(board)
				if dice[0] != 0 {
					lastRoll = dice

					advice, firstLine := Advise(board, dice, 2-j)
					fmt.Print(chalk.CyanLight())
					fmt.Print("Ranked computer generated advice...")
					fmt.Println(chalk.GreenLight())
					fmt.Print(advice)
					fmt.Println(chalk.Reset())

					board, dice = InterpretFinish(board, firstLine, dice)
				} else {
					break
				}
			}
			//fmt.Print("\033[H\033[2J")
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

	fmt.Print("\033[H\033[2J")
	fmt.Println(chalk.MagentaLight())
	fmt.Println("Congratulations! Great game!")
	fmt.Println(chalk.Reset())
	DisplayScoreBoard(board)
}
