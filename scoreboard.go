package main

import (
	"bufio"
	"fmt"
	"github.com/golang-demos/chalk"
	"os"
	"strings"
	"time"
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

// FindPossibleOptions Finds the playable options from the given dice and scoreboard and prompts the user to select one.
// Returns a new board
func FindPossibleOptions(board [13]ScoreItem,
	dice [5]int,
	fillInOption string,
	crossOut bool,
	crossOutOption string) [13]ScoreItem {
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
		yahtzeeValue := YahtzeeMultiplied

		possibilities = append(possibilities, ScoreItem{"Yahtzee", yahtzeeValue, 11})
	}
	if board[12].points == 0 { // This is the "Chance" option
		possibilities = append(possibilities, ScoreItem{"Chance", points, 12})
	}

	newBoard := board
	fmt.Print("Please type the name to fill in.")
	fmt.Println(chalk.YellowLight())

	for i := 0; i < len(possibilities); i++ {
		// Ex: Three-of-a-kind (30 points)
		fmt.Printf("%s (%d", possibilities[i].name, possibilities[i].points)
		if possibilities[i].points == 1 {
			fmt.Println(" point)")
		} else {
			fmt.Println(" points)")
		}
	}

	// They have to cross out
	if len(possibilities) == 0 {
		fmt.Println(chalk.RedLight())
		fmt.Printf("You have to cross out")
		fmt.Println(chalk.Reset())
		newBoard = CrossOut(newBoard, crossOutOption)
		return newBoard
	}

	fmt.Println(chalk.RedLight())
	fmt.Printf("Cross out")
	fmt.Println(chalk.Reset())

	badName := true
	scanner := bufio.NewScanner(os.Stdin)

	for badName {
		fmt.Print("Name > ")
		input := ""
		if fillInOption == "" && crossOut == false {
			if !scanner.Scan() {
				continue
			}
			input = scanner.Text()
		} else if crossOut == true {
			fmt.Println("cross out")
		} else {
			input = fillInOption
			fmt.Println(fillInOption)
		}

		if strings.ToLower(input) == "cross out" || crossOut == true {
			newBoard = CrossOut(newBoard, crossOutOption)
			break
		}

		for i := 0; i < len(possibilities); i++ {
			if strings.ToLower(possibilities[i].name) == strings.ToLower(input) {
				for j := 0; j < len(newBoard); j++ {
					if newBoard[j].id == possibilities[i].id {
						newBoard[j].points = possibilities[i].points

						if possibilities[i].id == 11 { // Add to Yahtzee  global variable
							YahtzeeValue = YahtzeeValue + YahtzeeMultiplied
							YahtzeeMultiplied += 50
							newBoard[j].points = YahtzeeValue
						}
					}
				}

				badName = false
				fmt.Println("Filled in", possibilities[i].name)
				break
			}
		}

		if badName {
			fmt.Println(chalk.RedLight())
			fmt.Print("Invalid name!")
			fmt.Println(chalk.Reset())
			fmt.Println()
			time.Sleep(2 * time.Second)
		}
	}

	return newBoard
}

func CrossOut(board [13]ScoreItem, crossOutOption string) [13]ScoreItem {
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
	scanner := bufio.NewScanner(os.Stdin)
	for badName {
		fmt.Print("Name > ")

		input := ""
		if crossOutOption == "" {
			if !scanner.Scan() {
				continue
			}
			input = scanner.Text()
		} else {
			input = crossOutOption
			fmt.Println(crossOutOption)
		}

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
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Println("Crossed out", name)
	return newBoard
}

func CalculateTargets(board [13]ScoreItem) {
	upperPoints := CalculateUpperScore(board)
	turnsLeft := ItemsNeededLeft(board)

	fmt.Println(chalk.CyanLight())
	fmt.Print("Personalized goals:")
	fmt.Println(chalk.YellowLight())

	goals := ""
	upperPossibleGoals := ""
	upperBackupGoals := ""

	variation := 0

	if board[0].points == 0 { // 1's
		variation += 3
		upperPossibleGoals += "You should use your 1's soon. Leave them as backup.\n"
		upperBackupGoals += "You should use your 1's soon. Leave them as backup.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 1's.\n")
	}
	if board[1].points == 0 { // 2's
		variation += 6
		upperPossibleGoals += "You should use your 2's soon. Leave them as backup.\n"
		upperBackupGoals += "You should use your 2's soon. Leave them as backup.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 2's.\n")
	}
	if board[2].points == 0 { // 3's
		variation += 9
		upperPossibleGoals += "You're close to the bonus. Try and get a lot of 3's.\n"
		upperBackupGoals += "Try for a bunch of 3's soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 3's.\n")
	}
	if board[3].points == 0 { // 4's
		variation += 12
		upperPossibleGoals += "You're close to the bonus. Try and get a lot of 4's.\n"
		upperBackupGoals += "Try for a bunch of 4's soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 4's.\n")
	}
	if board[4].points == 0 { // 5's
		variation += 15
		upperPossibleGoals += "You're close to the bonus. Try and get a lot of 5's.\n"
		upperBackupGoals += "Try for a bunch of 5's soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 5's.\n")
	}
	if board[5].points == 0 { // 6's
		variation += 18
		upperPossibleGoals += "You're close to the bonus. Try and get a lot of 6's.\n"
		upperBackupGoals += "Try for a bunch of 6's soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Roll for 6's.\n")
	}

	if 63-variation <= upperPoints && upperPoints < 63 && turnsLeft < 8 {
		goals += upperPossibleGoals
	} else if turnsLeft < 6 || (board[6].points != 0 &&
		board[9].points != 0 &&
		board[7].points != 0 &&
		board[8].points != 0 &&
		board[10].points != 0) {
		goals += upperBackupGoals
	}

	// Three-of-a-kind
	if turnsLeft > 2 && turnsLeft < 5 && board[6].points == 0 {
		goals += "You'll need a three-of-a-kind soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a three-of-a-kind.\n")
	} else if (turnsLeft <= 2 ||
		(board[9].points != 0 && board[7].points != 0 && board[8].points != 0 && board[10].points != 0)) &&
		board[6].points == 0 {
		goals += "Target a three-of-a-kind soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a three-of-a-kind.\n")
	}

	if board[10].points == 0 { // Large straight
		goals += "Target a large straight soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a large straight.\n")
	}
	if board[8].points == 0 { // Full house
		goals += "Target a full house soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a full house.\n")
	}
	if board[7].points == 0 && (turnsLeft < 6 || (board[8].points != 0 && board[10].points != 0)) { // Four-of-a-kind
		goals += "Target a four-of-a-kind soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a four-of-a-kind.\n")
	}

	// Small straight
	if turnsLeft > 3 && turnsLeft < 6 && board[9].points == 0 {
		goals += "You'll need a small straight soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a small straight.\n")
	} else if (turnsLeft <= 3 || (board[7].points != 0 && board[8].points != 0 && board[10].points != 0)) &&
		board[9].points == 0 {
		goals += "You should target a small straight soon.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a small straight.\n")
	}

	if board[11].points > 0 {
		goals += "Get a Yahtzee if you can. It'll be worth more.\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a Yahtzee!\n")
	}
	if board[12].points == 0 && turnsLeft < 7 {
		goals += "You still have your chance left as backup!\n"
		EmergencyAdvice = append(EmergencyAdvice, "Re-roll low numbers to get a good chance.\n")
	}

	if goals == "" {
		goals += "Try for a Yahtzee!\n"
		EmergencyAdvice = append(EmergencyAdvice, "Go for a Yahtzee!\n")
	}

	fmt.Print(goals)
	fmt.Println(chalk.Reset())
}

// CalculateTotalScore takes the board and returns the point values. Ignores -1 (crossed out) values
func CalculateTotalScore(board [13]ScoreItem) (score int) {
	for i := 0; i < len(board); i++ {
		if board[i].points != -1 {
			score += board[i].points
		}
	}

	// If they get a score of 63 or higher on their upper hand it's a +35 bonus
	if CalculateUpperScore(board) >= 63 {
		score += 35
	}

	return
}

func CalculateUpperScore(board [13]ScoreItem) int {
	return board[0].points + board[1].points + board[2].points + board[3].points + board[4].points + board[5].points
}

func DisplayScoreBoard(board [13]ScoreItem) {
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

	upperPoints := CalculateUpperScore(board)
	if upperPoints >= 63 {
		upperPoints += 35
	}

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

	points := CalculateTotalScore(board)

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
