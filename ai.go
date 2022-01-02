package main

import (
	"strconv"
	"strings"
)

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

func ItemsNeededLeft(board [13]ScoreItem) (total int) {
	for i := 0; i < len(board); i++ {
		if board[i].points == 0 {
			total++
		}
	}

	return
}

// Advise Returns advice and then the first item (for the computer).
// Board numbers:
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
func Advise(board [13]ScoreItem, dice [5]int, rollsLeft int) (string, string) {
	var advice string
	var adviceNum int

	points := 0
	upperPoints := CalculateUpperScore(board)
	for i := 0; i < len(dice); i++ {
		points += dice[i]
	}

	m := DiceMap(dice)
	// Advice about the Yahtzee
	if CalculateYahtzee(m) == 1 && board[12].points == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the Yahtzee and stop rolling.\n"
	}

	// Advice about locking in the 35 bonus points
	used1 := false
	used2 := false
	used3 := false
	used4 := false
	used5 := false
	used6 := false
	if upperStillOpen(board) && upperPoints < 63 {
		// We calculate how close they are to 63 points.
		// If this sends them over, make that the default priority
		// after a Yahtzee

		// Format: if they have that dice & upper points + that dice >= 63 & they haven't filled it in yet
		if m[1] > 0 && upperPoints+m[1] >= 61 && board[0].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 1's.\n"
				ActionString = "roll-non-1"
				used1 = true
			} else if upperPoints+m[1] >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 1's and get the bonus 35 points.\n"
				ActionString = "take-1"
				used1 = true
			}

		}
		if m[2] > 0 && upperPoints+(m[2]*2) >= 59 && board[1].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 2's.\n"
				used2 = true
				ActionString = "roll-non-2"
			} else if upperPoints+(m[2]*2) >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 2's and get the bonus 35 points.\n"
				used2 = true
				ActionString = "take-2"
			}
		}
		if m[3] > 0 && upperPoints+(m[3]*3) >= 57 && board[2].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 3's.\n"
				used3 = true
				ActionString = "roll-non-3"
			} else if upperPoints+(m[3]*3) >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 3's and get the bonus 35 points.\n"
				used3 = true
				ActionString = "take-3"
			}
		}
		if m[4] > 0 && upperPoints+(m[4]*4) >= 55 && board[3].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 4's.\n"
				used4 = true
				ActionString = "roll-non-4"
			} else if upperPoints+(m[4]*4) >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 4's and get the bonus 35 points.\n"
				used4 = true
				ActionString = "take-4"
			}
		}
		if m[5] > 0 && upperPoints+(m[5]*5) >= 63 && board[4].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 5's.\n"
				used5 = true
				ActionString = "roll-non-5"
			} else if upperPoints+(m[5]*5) >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 5's and get the bonus 35 points.\n"
				used5 = true
				ActionString = "take-5"
			}
		}
		if m[6] > 0 && upperPoints+(m[6]*6) >= 63 && board[5].points == 0 {
			adviceNum++
			if rollsLeft != 0 {
				advice += strconv.Itoa(adviceNum) + ". You're " +
					strconv.Itoa(63-upperPoints) + " points from the bonus. Get more 6's.\n"
				used6 = true
				ActionString = "roll-non-6"
			} else if upperPoints+(m[6]*6) >= 63 {
				advice += strconv.Itoa(adviceNum) + ". Take your 6's and get the bonus 35 points.\n"
				used6 = true
				ActionString = "take-6"
			}
		}
	}

	// Next we suggest taking things we already have
	if board[10].points == 0 && CalculateLargeStraight(dice) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the large straight and stop rolling.\n"
		ActionString = "take-large"
	}
	if board[8].points == 0 && CalculateFullHouse(m) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the Full house and stop rolling.\n"
		ActionString = "take-1's"
	}

	if CalculateFourKind(m) > 0.1 && CalculateFourKind(m) != 1 && board[7].points == 0 && rollsLeft != 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Go for a four-of-a-kind.\n"
	}
	if CalculateThreeKind(m) == 1 && rollsLeft == 2 && board[11].points != -1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Go for a Yahtzee.\n"
	}

	largePicked := false
	if CalculateSmallStraight(dice) == 1 && CalculateLargeStraight(dice) != 1 && board[10].points == 0 &&
		rollsLeft != 0 {
		adviceNum++
		largePicked = true
		advice += strconv.Itoa(adviceNum) + ". Go for a large straight.\n"
	}

	if board[9].points == 0 && CalculateSmallStraight(dice) == 1 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the small straight.\n"
	}

	// Next we recommend taking the upper hand if they get three or more
	// unless they're already filled in because you're more likely to find a
	// four 5's useful in your 5's instead of as a four-of-a-kind
	threePicked := false
	fourPicked := false
	if rollsLeft == 0 {
		if board[3].points == 0 && m[4] >= 4 && !used4 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 4's.\n"
		}
		if board[4].points == 0 && m[5] >= 4 && !used5 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 5's.\n"
		}
		if board[5].points == 0 && m[6] >= 4 && !used6 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 6's.\n"
		}
		if board[7].points == 0 && CalculateFourKind(m) == 1 && points > 15 {
			adviceNum++
			fourPicked = true
			advice += strconv.Itoa(adviceNum) + ". Take the four-of-a-kind.\n"
		}
		if board[6].points == 0 && CalculateThreeKind(m) == 1 && points > 20 {
			adviceNum++
			threePicked = true
			advice += strconv.Itoa(adviceNum) + ". Take the three-of-a-kind.\n"
		}
		if board[3].points == 0 && m[4] >= 3 && m[4] < 4 && !used4 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 4's.\n"
		}
		if board[4].points == 0 && m[5] >= 3 && m[5] < 4 && !used4 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 5's.\n"
		}
		if board[5].points == 0 && m[6] >= 3 && m[6] < 4 && !used6 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Take your 6's.\n"
		}
	}

	if ItemsNeededLeft(board) < 6 || rollsLeft == 2 {
		if CalculateLargeStraight(dice) > 0.1 && CalculateLargeStraight(dice) != 1 && board[10].points == 0 &&
			rollsLeft != 0 && !largePicked {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Go for a large straight.\n"
		}
		if CalculateFullHouse(m) > 0.1 && CalculateFullHouse(m) != 1 && board[8].points == 0 && rollsLeft != 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Go for a full house.\n"
		}
	}

	if rollsLeft > 0 {
		if m[6] >= 2 {
			if board[5].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 6's and keep rolling.\n"
			}
		}
		if m[5] >= 2 {
			if board[4].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 5's and keep rolling.\n"
			}
		}
		if m[4] >= 2 {
			if board[3].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 4's and keep rolling.\n"
			}
		}
		if m[3] >= 2 {
			if board[2].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 3's and keep rolling.\n"
			}
		}
		if m[2] >= 2 {
			if board[1].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 2's and keep rolling.\n"
			}
		}
		if m[1] >= 2 {
			if board[0].points == 0 || board[6].points == 0 || board[7].points == 0 || board[12].points == 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Keep your 1's and keep rolling.\n"
			}
		}
		if CalculateSmallStraight(dice) > 0.1 && CalculateSmallStraight(dice) != 1 && board[9].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Go for a small straight.\n"
		}
		if CalculateYahtzee(m) > 0.1 && ItemsNeededLeft(board) >= 6 && board[11].points != -1 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Go for a Yahtzee.\n"
		}
		if ItemsNeededLeft(board) < 5 && board[12].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Re-roll low numbers to get a good chance.\n"
		}
	}

	// Always go for chance and small numbers as a last resort
	if board[12].points == 0 && points > 23 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Use your chance.\n"
	} else if board[12].points == 0 && adviceNum <= 1 && ItemsNeededLeft(board) < 5 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Use your chance.\n"
	}
	if board[0].points == 0 && m[1] != 0 && adviceNum <= 1 && !used1 && rollsLeft == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take your 1's.\n"
	}
	if board[1].points == 0 && m[2] != 0 && adviceNum <= 1 && !used2 && rollsLeft == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take your 2's.\n"
	}
	if board[2].points == 0 && m[3] != 0 && adviceNum <= 1 && !used3 && rollsLeft == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take your 3's.\n"
	}

	if board[7].points == 0 && CalculateFourKind(m) == 1 && !fourPicked && rollsLeft == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the four-of-a-kind.\n"
	}
	if board[6].points == 0 && CalculateThreeKind(m) == 1 && !threePicked && rollsLeft == 0 {
		adviceNum++
		advice += strconv.Itoa(adviceNum) + ". Take the three-of-a-kind.\n"
	}

	if adviceNum == 0 && ItemsNeededLeft(board) == 1 && rollsLeft == 0 {
		for i := 0; i < 7; i++ {
			if board[i].points == 0 && m[i+1] > 0 {
				adviceNum++
				advice += strconv.Itoa(adviceNum) + ". Take your " + strconv.Itoa(i+1) + "'s.\n"
			}
		}

	}

	if adviceNum == 0 && rollsLeft != 0 {
		for i := 0; i < len(EmergencyAdvice); i++ {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". " + EmergencyAdvice[i]
		}
	}

	// If we get here and there's still nothing cross out in this order
	// There's a few duplicates because there's lots of conditions (like how few turns left,
	// how close to certain milestones, etc.)
	if adviceNum == 0 {
		if board[11].points == 0 && ItemsNeededLeft(board) < 6 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your Yahtzee.\n"
		} else if board[7].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your four-of-a-kind.\n"
		} else if board[6].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your three-of-a-kind.\n"
		} else if board[0].points == 0 && upperPoints-3 < 63 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 1's.\n"
		} else if board[1].points == 0 && upperPoints-6 < 63 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 2's.\n"
		} else if board[8].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your full house.\n"
		} else if board[10].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your large straight.\n"
		} else if board[2].points == 0 && upperPoints-9 < 63 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 3's.\n"
		} else if board[3].points == 0 && upperPoints-12 < 63 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 4's.\n"
		} else if board[11].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your Yahtzee.\n"
		} else if board[7].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your four-of-a-kind.\n"
		} else if board[6].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your three-of-a-kind.\n"
		} else if board[0].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 1's.\n"
		} else if board[1].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 2's.\n"
		} else if board[2].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 3's.\n"
		} else if board[3].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 4's.\n"
		} else if board[4].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 5's.\n"
		} else if board[5].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your 6's.\n"
		} else if board[9].points == 0 {
			adviceNum++
			advice += strconv.Itoa(adviceNum) + ". Cross out your small straight.\n"
		}
	}

	var firstLine string
	for _, line := range strings.Split(strings.TrimSuffix(advice, "\n"), "\n") {
		firstLine = line
		break
	}

	return advice, firstLine
}
