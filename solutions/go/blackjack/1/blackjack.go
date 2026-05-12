package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var i int
	switch card {
	case "ace":
		i = 11
	case "two":
		i = 2
	case "three":
		i = 3
	case "four":
		i = 4
	case "five":
		i = 5
	case "six":
		i = 6
	case "seven":
		i = 7
	case "eight":
		i = 8
	case "nine":
		i = 9
	case "ten", "jack", "queen", "king":
		i = 10
	default:
		i = 0
	}

	return i

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, d1 := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)

	sum := c1 + c2

	var i string
	switch {
	// If you have a pair of aces you must always split them.
	case c1 == 11 && c2 == 11:
		i = "P"

	// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	case sum == 21:
		if d1 != 11 && d1 != 10 {
			i = "W"
		} else {
			i = "S"
		}

	// If your cards sum up to a value within the range [17, 20] you should always stand.
	case sum >= 17 && sum <= 20:
		i = "S"

	// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	case sum >= 12 && sum <= 16:
		if d1 >= 7 {
			i = "H"
		} else {
			i = "S"
		}

	//If your cards sum up to 11 or lower you should always hit.
	case sum <= 11:
		i = "H"
	}

	return i

}
