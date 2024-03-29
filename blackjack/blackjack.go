package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	var cardValue int
	switch card {
	case "ace":
		cardValue = 11
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten", "jack", "queen", "king":
		cardValue = 10
	case "other":
		cardValue = 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var cardTotal int = ParseCard(card1) + ParseCard(card2)
	var selectedOption string

	switch {
	case card1 == "ace" && card2 == "ace":
		selectedOption = "P"

	case cardTotal == 21 && dealerCard != "ace" && ParseCard(dealerCard) != 10:
		selectedOption = "W"
	case cardTotal == 21 && (dealerCard == "ace" || ParseCard(dealerCard) == 10):
		selectedOption = "S"

	case cardTotal >= 17 && cardTotal <= 20:
		selectedOption = "S"

	case cardTotal >= 12 && cardTotal <= 16 && ParseCard(dealerCard) < 7:
		selectedOption = "S"
	case cardTotal >= 12 && cardTotal <= 16 && ParseCard(dealerCard) >= 7:
		selectedOption = "H"

	case cardTotal <= 11:
		selectedOption = "H"
	}

	return selectedOption
}
