// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("CyanCoding's Yahtzee Hack!")
	totalSuccesses := 0.0
	numberOfTrials := 3.0
	phase1 := Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)
	totalSuccesses = 1.0
	phase2 := Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)
	totalSuccesses = 2.0
	phase3 := Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)

	value := Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)
	totalSuccesses = 3.0
	value += Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)

	totalSuccesses = 1.0
	numberOfTrials = 2.0
	value2 := Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)
	totalSuccesses = 2.0
	value2 += Factorial(numberOfTrials) / (Factorial(numberOfTrials-totalSuccesses) * Factorial(totalSuccesses)) * math.Pow(1.0/6.0, totalSuccesses) * math.Pow(5.0/6.0, numberOfTrials-totalSuccesses)

	percentage := ((phase1 * value) + (phase2 * value2)) + phase3

	fmt.Println(percentage)

	score := 0

	// Infinite loop until the game ends
	for i := 1; i <= 13; i++ {
		fmt.Println()
		fmt.Println("------------")
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
			fmt.Print("\033[H\033[2J")
			for i := 0; i < 3; i++ { // Up to three rolls per turn
				dice := InputDice()
				if dice[0] != 0 {
					CalculateLowerHand(dice)
				} else {
					break
				}
			}
		} else if input == 3 { // Calculate score

		} else {
			fmt.Println("Invalid action. Try again.")
			i-- // We don't want to waste a turn on a bad action
			continue
		}

	}
}
