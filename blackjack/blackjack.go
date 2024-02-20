package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "eight":
		return 8
	case "two":
		return 2
	case "nine":
		return 9
	case "three":
		return 3
	case "ten":
		return 10
	case "four":
		return 4
	case "jack":
		return 10
	case "five":
		return 5
	case "queen":
		return 10
	case "six":
		return 6
	case "king":
		return 10
	case "seven":
		return 7
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	} else if (sum >= 17) && (sum <= 20) {
		return "S"
	} else if (sum >= 12) && (sum <= 16) {
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	} else if sum <= 11 {
		return "H"
	} else {
		if (ParseCard(card1) == 10 && ParseCard(dealerCard) < 10) || (ParseCard(card2) == 10 && ParseCard(dealerCard) < 10) {
			return "W"
		} else {
			return "S"
		}
	}
}
