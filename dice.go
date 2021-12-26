package main

import (
	"fmt"
	"github.com/fatih/color"
)

func InputDice() (dice [5]int) {
	var input string
	for len(input) != 5 || input == "0" {
		fmt.Print("Please enter your dice ('34531') or '0' to quit > ")
		_, _ = fmt.Scanln(&input)

		if input == "0" {
			dice[0] = 0
			return
		} else if len(input) != 5 {
			color.Set(color.FgHiRed)
			fmt.Println("You should have 5 dice!")
			color.Set(color.FgHiWhite)
		}
	}

	// Add to dice array
	for i := 0; i < 5; i++ {
		dice[i] = int(input[i] - '0')
	}

	return
}

func calculateThreeKind(dice [5]int) (probability float32) {
func factorial(n int64) (value int64) {
	if n > 0 {
		value = n * Factorial(n-1)
		return
	}

	return 1
}
	return
}

func CalculateLowerHand(dice [5]int) {
	//var threeKind, fourKind, fullHouse, smallStraight, largeStraight, yahtzee float32
}
