package main

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"strconv"
	"strings"
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

func Factorial(n float64) (value float64) {
	if n > 0 {
		value = n * Factorial(n-1)
		return
	}

	return 1
}

func calculateProbability(rolling float64, needed float64) (probability float64) {
	if needed > rolling {
		return 0
	}
	// Combinations = n!/(r!(n - r)!)
	combinations := Factorial(rolling) / (Factorial(needed) * Factorial(rolling-needed))

	// Probability = combinations * (1/6)^needed * (1 - (1/6))^(n - r)
	probability = combinations * math.Pow(1.0/6.0, needed) * math.Pow(1-(1.0/6.0), rolling-needed)

	return
}

func calculateThreeKind(m map[int]int) (probability float64) {
	for _, value := range m {
		if value >= 3 {
			return 1
		}

		temp := calculateProbability(float64(5-value), float64(3-value))

		// This way we always get the highest probability returned
		if temp > probability {
			probability = temp
		}
	}

	return
}

func calculateFourKind(m map[int]int) (probability float64) {
	for _, value := range m {
		if value >= 4 {
			return 1
		}

		temp := calculateProbability(float64(5-value), float64(4-value))

		// This way we always get the highest probability returned
		if temp > probability {
			probability = temp
		}
	}

	return
}

func calculateFullHouse(m map[int]int) (probability float64) {
	haveThree := 0
	haveTwo := 0
	for _, value := range m {
		if value == 2 {
			haveTwo++
		} else if value == 3 {
			haveThree++
		}
	}

	if haveThree == 1 && haveTwo == 1 {
		return 1 // They already have a full house
	} else if haveTwo == 2 { // They have two pairs so they only need one extra card
		probability = calculateProbability(1.0, 1.0) // This evaluates to 16.7%
	} else if haveTwo == 1 { // So they need a pair to make a loose 1 into 3
		probability = calculateProbability(3.0, 2.0)
	} else if haveThree == 1 { // They just need one match
		probability = calculateProbability(1.0, 1.0) // This is also 16.7%
	} else { // They have nothing good
		probability = calculateProbability(4.0, 4.0)
	}
	return
}

func calculateSmallStraight(dice [5]int) (probability float64) {
	possibility1 := "1234"
	possibility2 := "2345"
	possibility3 := "3456"

	for i := 0; i < 5; i++ {
		character := strconv.Itoa(dice[i])
		if strings.Contains(possibility1, character) {
			possibility1 = strings.Replace(possibility1, character, "", -1)
		}
		if strings.Contains(possibility2, character) {
			possibility2 = strings.Replace(possibility2, character, "", -1)
		}
		if strings.Contains(possibility3, character) {
			possibility3 = strings.Replace(possibility3, character, "", -1)
		}
	}

	if len(possibility1) == 0 || len(possibility2) == 0 || len(possibility3) == 0 {
		// They already have a small straight!
		return 1
	}

	// Get the lowest number for use in the prediction
	// Since we would go for a straight we're 1 away from
	// over one we're 3 away from
	numToUse := len(possibility1)
	if len(possibility2) < numToUse {
		numToUse = len(possibility2)
	}
	if len(possibility3) < numToUse {
		numToUse = len(possibility3)
	}

	// 5.0 (dice) - (4.0 (length of small straight) - num to use
	// Since if we have 1 number we need, that means we have three good dice (4 - 1), so we re-roll 2.
	probability = calculateProbability(5.0-(4.0-float64(numToUse)), float64(numToUse))

	return
}

func calculateLargeStraight(dice [5]int) (probability float64) {
	possibility1 := "12345"
	possibility2 := "23456"

	for i := 0; i < 5; i++ {
		character := strconv.Itoa(dice[i])
		if strings.Contains(possibility1, character) {
			possibility1 = strings.Replace(possibility1, character, "", -1)
		}
		if strings.Contains(possibility2, character) {
			possibility2 = strings.Replace(possibility2, character, "", -1)
		}
	}

	if len(possibility1) == 0 || len(possibility2) == 0 {
		// They already have a large straight!
		return 1
	}

	// Get the lowest number for use in the prediction
	// Since we would go for a straight we're 1 away from
	// over one we're 3 away from
	numToUse := len(possibility1)
	if len(possibility2) < numToUse {
		numToUse = len(possibility2)
	}

	probability = calculateProbability(float64(numToUse), float64(numToUse))

	return
}

func calculateYahtzee(m map[int]int) (probability float64) {
	for _, value := range m {
		if value == 5 {
			return 1
		}

		temp := calculateProbability(float64(5-value), float64(5-value))

		// This way we always get the highest probability returned
		if temp > probability {
			probability = temp
		}
	}

	return
}

func diceMap(dice [5]int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < 5; i++ {
		m[dice[i]]++
	}

	return m
}

func CalculateLowerHand(dice [5]int) {
	var threeKind, fourKind, fullHouse, smallStraight, largeStraight, yahtzee float64

	m := diceMap(dice)

	// Sort values into a map so that we know how many of each number there is
	points := 0
	for i := 0; i < 5; i++ {
		points += dice[i]
	}

	threeKind = calculateThreeKind(m) * 100
	fourKind = calculateFourKind(m) * 100
	fullHouse = calculateFullHouse(m) * 100
	smallStraight = calculateSmallStraight(dice) * 100
	largeStraight = calculateLargeStraight(dice) * 100
	yahtzee = calculateYahtzee(m) * 100

	// Print results
	fmt.Println()
	color.Set(color.FgCyan)
	fmt.Println("In your next roll you might get...")
	color.Set(color.FgHiGreen)

	if yahtzee == 100 {
		color.Set(color.FgHiMagenta)
		fmt.Println("----- YAHTZEE (50 points) ✓ -----")
		color.Set(color.FgHiGreen)
		fullHouse = 100 // Joker rule
	}
	if threeKind == 100 {
		fmt.Println("Three of a kind (" + strconv.Itoa(points) + " points): ✓")
	}
	if fourKind == 100 {
		fmt.Println("Four of a kind (" + strconv.Itoa(points) + " points): ✓")
	}
	if fullHouse == 100 {
		fmt.Println("Full house (25 points): ✓")
	}
	if smallStraight == 100 {
		fmt.Println("Small straight (30 points): ✓")
	}
	if largeStraight == 100 {
		fmt.Println("Large straight (40 points): ✓")
	}

	color.Set(color.FgHiYellow)

	if threeKind != 100 {
		fmt.Println("Three of a kind: " + strconv.FormatFloat(threeKind, 'f', 2, 64) + "%")
	}
	if fourKind != 100 {
		fmt.Println("Four of a kind: " + strconv.FormatFloat(fourKind, 'f', 2, 64) + "%")
	}
	if fullHouse != 100 {
		fmt.Println("Full house: " + strconv.FormatFloat(fullHouse, 'f', 2, 64) + "%")
	}
	if smallStraight != 100 {
		fmt.Println("Small straight: " + strconv.FormatFloat(smallStraight, 'f', 2, 64) + "%")
	}
	if largeStraight != 100 {
		fmt.Println("Large straight: " + strconv.FormatFloat(largeStraight, 'f', 2, 64) + "%")
	}
	if yahtzee != 100 {
		fmt.Println("Yahtzee: " + strconv.FormatFloat(yahtzee, 'f', 2, 64) + "%")
	} else if yahtzee == 100 {
		// We have to reset dice
		for i := 1; i < 6; i++ {
			dice[i-1] = i
		}
		m = diceMap(dice)

		yahtzee = calculateYahtzee(m) * 100
		fmt.Println("Another Yahtzee: " + strconv.FormatFloat(yahtzee, 'f', 2, 64) + "%")
	}

	color.Set(color.FgHiWhite)
	fmt.Println()
}
