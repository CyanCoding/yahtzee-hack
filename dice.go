package main

import (
	"fmt"
)

func InputDice() (dice [5]int) {
	var input string
	for len(input) != 5 {
		fmt.Print("Please enter your dice ('34531') > ")
		_, _ = fmt.Scanln(&input)

		if len(input) != 5 {
			fmt.Println("You should have 5 dice!")
		}
	}

	// Add to dice array
	for i := 0; i < 5; i++ {
		dice[i] = int(input[i] - '0')
	}

	return
}

func CalculateProbability(dice [5]int) {

}
