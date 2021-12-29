package main

import (
	"fmt"
	"github.com/golang-demos/chalk"
	"strings"
)

// ScoreItem holds data on a scoreboard item's name, current point value, and ID
type ScoreItem struct {
	name   string
	points int
	id     int
}

// GenerateBoard creates blank (0's) data for each playable option
func GenerateBoard() [13]ScoreItem {
	board := [13]ScoreItem{}

	board[0] = ScoreItem{"1's", 0, 0}
	board[1] = ScoreItem{"2's", 0, 1}
	board[2] = ScoreItem{"3's", 0, 2}
	board[3] = ScoreItem{"4's", 0, 3}
	board[4] = ScoreItem{"5's", 0, 4}
	board[5] = ScoreItem{"6's", 0, 5}
	board[6] = ScoreItem{"Three-of-a-kind", 0, 6}
	board[7] = ScoreItem{"Four-of-a-kind", 0, 7}
	board[8] = ScoreItem{"Full house", 0, 8}
	board[9] = ScoreItem{"Small straight", 0, 9}
	board[10] = ScoreItem{"Large straight", 0, 10}
	board[11] = ScoreItem{"Yahtzee", 0, 11}
	board[12] = ScoreItem{"Chance", 0, 12}

	return board
}

// FindPossibleOptions Finds the playable options from the given dice and scoreboard and prints them out.
// Returns the vector of playable options with "Cross out" in the last position.
func FindPossibleOptions(board [13]ScoreItem, dice [5]int) []ScoreItem {
	points := 0
	for i := 0; i < 5; i++ {
		points += dice[i]
	}

	m := DiceMap(dice)
	possibilities := make([]ScoreItem, 0)

	// Calculate upper hand
	// Note: it's m[1] instead of m[0] because 1 references the number, not the position
	if m[1] > 0 && board[0].points == 0 {
		possibilities = append(possibilities, ScoreItem{"1's", m[1], 0})
	}

	if m[2] > 0 && board[1].points == 0 {
		possibilities = append(possibilities, ScoreItem{"2's", m[2] * 2, 1})
	}

	if m[3] > 0 && board[2].points == 0 {
		possibilities = append(possibilities, ScoreItem{"3's", m[3] * 3, 2})
	}

	if m[4] > 0 && board[3].points == 0 {
		possibilities = append(possibilities, ScoreItem{"4's", m[4] * 4, 3})
	}

	if m[5] > 0 && board[4].points == 0 {
		possibilities = append(possibilities, ScoreItem{"5's", m[5] * 5, 4})
	}

	if m[6] > 0 && board[5].points == 0 {
		possibilities = append(possibilities, ScoreItem{"6's", m[6] * 6, 5})
	}

	// Calculate lower hand
	if CalculateThreeKind(m) == 1 && board[6].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Three-of-a-kind", points, 6})
	}

	if CalculateFourKind(m) == 1 && board[7].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Four-of-a-kind", points, 7})
	}

	if CalculateFullHouse(m) == 1 && board[8].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Full house", 25, 8})
	}

	if CalculateSmallStraight(dice) == 1 && board[9].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Small straight", 30, 9})
	}

	if CalculateLargeStraight(dice) == 1 && board[10].points == 0 {
		possibilities = append(possibilities, ScoreItem{"Large straight", 40, 10})
	}
	if CalculateYahtzee(m) == 1 {
		// A Yahtzee is worth +50 points each time.
		// The first is worth 50, then 100, then 150, etc.
		yahtzeeValue := board[11].points + 50
		possibilities = append(possibilities, ScoreItem{"Yahtzee", yahtzeeValue, 11})
	}
	if board[12].points == 0 { // This is the "Chance" option
		possibilities = append(possibilities, ScoreItem{"Chance", points, 12})
	}

	var i int
	for i = 0; i < len(possibilities); i++ {
		// Ex: 4. Three-of-a-kind (30 points)
		fmt.Printf("%d. %s (%d", i+1, possibilities[i].name, possibilities[i].points)
		if possibilities[i].points == 1 {
			fmt.Println(" point)")
		} else {
			fmt.Println(" points)")
		}
	}
	fmt.Println(chalk.RedLight())
	fmt.Printf("%d. Cross out", i+1)
	fmt.Println(chalk.Reset())

	possibilities = append(possibilities, ScoreItem{"Cross out", 0, 13})
	return possibilities
}

func CrossOut(board [13]ScoreItem) [13]ScoreItem {
	newBoard := board
	fmt.Print("\033[H\033[2J")
	fmt.Print("Please type in the name to cross it out.")
	fmt.Println(chalk.YellowLight())

	// Counter represents the 1., 2., 3., 4., etc. counter used to display
	// an orderly list to the user. We can't use board[i].id or i because
	// they'll skip over ones we can't cross out
	counter := 1
	for i := 0; i < len(board); i++ {
		if board[i].points == 0 {
			// Ex: 4. Three-of-a-kind
			fmt.Printf("%s\n", board[i].name)
			counter++
		}
	}
	fmt.Println(chalk.Reset())

	var name string
	badName := true
	for badName {
		var input string
		fmt.Print("Name > ")
		_, _ = fmt.Scanln(&input)
		for i := 0; i < len(newBoard); i++ {
			if strings.ToLower(newBoard[i].name) == strings.ToLower(input) {
				newBoard[i].points = -1
				name = newBoard[i].name
				badName = false
				break
			}
		}

		if badName {
			fmt.Println(chalk.RedLight())
			fmt.Print("Invalid name!")
			fmt.Println(chalk.Reset())
			fmt.Println()
		}
	}

	fmt.Println("Crossed out", name)
	return newBoard
}

// CalculateScore takes the board and returns the point values. Ignores -1 (crossed out) values
func CalculateScore(board [13]ScoreItem) (score int) {
	for i := 0; i < len(board); i++ {
		if board[i].points != -1 {
			score += board[i].points
		}
	}

	return
}

func DisplayScoreBoard(board [13]ScoreItem) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("==== SCORECARD ====")
	fmt.Println("-------------------")
	fmt.Println("|  Upper Section  |")
	fmt.Print("-------------------")

	for i := 0; i < 6; i++ {
		if board[i].points == -1 {
			fmt.Println(chalk.Strikethrough())
			fmt.Printf("| %s             : 0", board[i].name)
		} else {
			fmt.Println(chalk.Reset())
			fmt.Printf("| %s             : %d", board[i].name, board[i].points)
		}
	}
	fmt.Println(chalk.Reset())
	fmt.Println("-------------------")

	upperPoints := board[0].points + board[1].points + board[2].points + board[3].points + board[4].points + board[5].points
	if upperPoints < 10 {
		fmt.Printf("|    Points: %d    |", upperPoints)
	} else if upperPoints < 100 {
		fmt.Printf("|    Total: %d    |", upperPoints)
	} else {
		fmt.Printf("|   Points: %d   |", upperPoints)
	}

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()
	fmt.Println("-------------------")

	fmt.Println("|  Lower Section  |")
	fmt.Print("-------------------")

	if board[6].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Three-of-a-kind : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Three-of-a-kind : %d", board[6].points)
	}

	if board[7].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Four-of-a-kind  : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Four-of-a-kind  : %d", board[7].points)
	}

	if board[8].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Full house      : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Full house      : %d", board[8].points)
	}

	if board[9].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Small straight  : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Small straight  : %d", board[9].points)
	}

	if board[10].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Large straight  : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Large straight  : %d", board[10].points)
	}

	if board[11].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Yahtzee         : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Yahtzee         : %d", board[11].points)
	}

	if board[12].points == -1 {
		fmt.Println(chalk.Strikethrough())
		fmt.Printf("| Chance          : 0")
	} else {
		fmt.Println(chalk.Reset())
		fmt.Printf("| Chance          : %d", board[12].points)
	}

	fmt.Println(chalk.Reset())
	fmt.Println("-------------------")

	points := CalculateScore(board)

	fmt.Println(chalk.GreenLight())

	if points < 10 {
		fmt.Printf("==== POINTS: %d ====", points)
	} else if points < 100 {
		fmt.Printf("==== TOTAL: %d ====", points)
	} else if points < 1000 {
		fmt.Printf("=== POINTS: %d ===", points)
	} else if points < 10000 {
		fmt.Printf("=== TOTAL: %d ===", points)
	}
	fmt.Println(chalk.Reset())
}
