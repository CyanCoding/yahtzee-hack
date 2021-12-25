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

func calculateThreeKind(dice [5]int) (probability float32) {
	return
}

func CalculateLowerHand(dice [5]int) {
	//var threeKind, fourKind, fullHouse, smallStraight, largeStraight, yahtzee float32
}
