package main

import (
	"fmt"
)

var dice [5]int

func InputDice() {
	var input string
	for len(input) != 5 {
		fmt.Print("Please enter your dice ('34531') > ")
		_, _ = fmt.Scanln(&input)

		if len(input) != 5 {
			fmt.Println("You should have 5 dice!")
		}
	}

	// Add to global dice array
	for i := 0; i < 5; i++ {
		dice[i] = int(input[i] - '0')
	}
}
