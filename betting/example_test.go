package betting_test

import (
	"fmt"

	"github.com/gustavooferreira/bfutils/betting"
)

// This example computes the bet type and size to be placed on the market in order to
// perform a "Green Book" operation.
func Example_a() {
	selection := betting.Selection{
		Bets: []betting.Bet{
			{Type: betting.BetType_Back, Odd: 4, Amount: 5},
			{Type: betting.BetType_Lay, Odd: 3, Amount: 5},
			{Type: betting.BetType_Back, Odd: 3.5, Amount: 10},
			{Type: betting.BetType_Lay, Odd: 3.2, Amount: 10},
		},
		CurrentBackOdd: 2.4,
		CurrentLayOdd:  2.42,
	}

	bet, err := betting.GreenBookSelection(selection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("In order to green book this selection, put a {%s} bet at {%.2f} for £%.2f.\n",
		bet.Type, bet.Odd, bet.Amount)

	fmt.Printf("P&L\n")
	fmt.Printf("---\n")
	fmt.Printf("If this selection wins:  £%.2f\n", bet.WinPL)
	fmt.Printf("If this selection loses: £%.2f\n", bet.LosePL)

	// Output:
	// In order to green book this selection, put a {Lay} bet at {2.42} for £3.31.
	// P&L
	// ---
	// If this selection wins:  £3.31
	// If this selection loses: £3.31
}
