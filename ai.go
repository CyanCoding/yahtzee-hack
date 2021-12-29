package main

import "strconv"

func upperStillOpen(board [13]ScoreItem) bool {
	if board[0].points == 0 ||
		board[1].points == 0 ||
		board[2].points == 0 ||
		board[3].points == 0 ||
		board[4].points == 0 ||
		board[5].points == 0 {
		return true
	}

	return false
}

// Advise board numbers:
/* board[0] = "1's"
 * board[1] = "2's"
 * board[2] = "3's"
 * board[3] = "4's"
 * board[4] = "5's"
 * board[5] = "6's"
 * board[6] = "Three-of-a-kind"
 * board[7] = "Four-of-a-kind"
 * board[8] = "Full house"
 * board[9] = "Small straight"
 * board[10] = "Large straight"
 * board[11] = "Yahtzee"
 * board[12] = "Chance"
 */
func Advise(board [13]ScoreItem, dice [5]int, rollsLeft int, turnsLeft int) {
	var advice string
	var adviceNum int

	m := DiceMap(dice)
	// Advice about the Yahtzee
	if CalculateYahtzee(m) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the Yahtzee.\n"
	}

	// Advice about locking in the 35 bonus points
	if upperStillOpen(board) {
		upperPoints := CalculateUpperScore(board)

		// We calculate how close they are to 63 points.
		// If this sends them over, make that the default priority
		// after a Yahtzee

		// Format: if they have that dice & upper points + that dice >= 63 & they haven't filled it in yet
		if m[1] > 0 && upperPoints+m[1] >= 63 && board[0].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 1's and get the bonus 35 points.\n"
		}
		if m[2] > 0 && upperPoints+m[2] >= 63 && board[1].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 2's and get the bonus 35 points.\n"
		}
		if m[3] > 0 && upperPoints+m[3] >= 63 && board[2].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 3's and get the bonus 35 points.\n"
		}
		if m[4] > 0 && upperPoints+m[4] >= 63 && board[3].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 4's and get the bonus 35 points.\n"
		}
		if m[5] > 0 && upperPoints+m[5] >= 63 && board[4].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 5's and get the bonus 35 points.\n"
		}
		if m[6] > 0 && upperPoints+m[6] >= 63 && board[5].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 6's and get the bonus 35 points.\n"
		}
	}

	// Next we suggest taking things we already have
	if board[10].points == 0 && CalculateLargeStraight(dice) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the large straight.\n"
	}
	if board[8].points == 0 && CalculateFullHouse(m) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the Full house.\n"
	}

	// Add in a few exceptions here, like if we had all our upper ones filled but not a small straight
	if !upperStillOpen(board) && board[9].points == 0 && CalculateSmallStraight(dice) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the small straight.\n"
	}
	
	// TODO: A few things to add: Next roll advice, advice based on how many turns are left, etc.


	if board[7].points == 0 && CalculateFourKind(m) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the four-of-a-kind.\n"
	}

}
