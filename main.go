// File created on Dec 25, 2021
// Author: CyanCoding, Camden Mac Leod
package main

import "fmt"

func main() {
	fmt.Println("CyanCoding's Yahtzee Hack!")

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

		if input == 1 {

		} else if input == 2 {

		} else if input == 3 {

		} else {
			fmt.Println("Invalid action. Try again.")
			i-- // We don't want to waste a turn on a bad action
			continue
		}

	}
}
