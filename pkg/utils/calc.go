package utils

// FreeBetDecimal returns the P&L multiplier factor
// Example: back:4 lay:2 the multiplier factor is 2, which means if you back at odd 4 with £10
// and lay at odd 2 with £10 you secure a free bet of 2 * £10 = £20
func FreeBetDecimal(oddBack float64, oddLay float64) float64 {
	return oddBack - oddLay
}

// FreeBetAmount returns the profit in case selection wins.
// Note that 'stake' is the backer's stake not the layer's liability
func FreeBetAmount(oddBack float64, oddLay float64, stake float64) float64 {
	return FreeBetDecimal(oddBack, oddLay) * stake
}

// GreenBookOpenBackAmount returns lay stake to greenbook.
func GreenBookOpenBackAmount(oddBack float64, stakeBack float64, oddLay float64) float64 {
	return (stakeBack * oddBack) / oddLay
}

// GreenBookOpenBackDecimal returns percentage of P&L
func GreenBookOpenBackDecimal(oddBack float64, oddLay float64) float64 {
	return oddBack/oddLay - 1
}

// GreenBookOpenBackAmountByPerc returns oddLay for a given perc P&L
// Note that when Backing, you cannot lose more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
// perc is a representation in decimal, meaning if you want to know at what LAY odd you should
// place a bet at in order to get 100% profit, then perc is == 1
func GreenBookOpenBackAmountByPerc(oddBack float64, perc float64) float64 {
	return oddBack / (perc + 1)
}

// GreenBookOpenBackAmount returns back stake to greenbook.
func GreenBookOpenLayAmount(oddLay float64, stakeLay float64, oddBack float64) float64 {
	return (stakeLay * oddLay) / oddBack
}

// GreenBookOpenLayDecimal returns percentage of P&L
func GreenBookOpenLayDecimal(oddLay float64, oddBack float64) float64 {
	return 1 - oddLay/oddBack
}

// GreenBookOpenLayAmountByPerc returns oddBack for a given perc P&L
// Note that when Laying, you cannot win more than 100% of your stake
// therefore feeding perc with a number less or equal to -1 is an error!
func GreenBookOpenLayAmountByPerc(oddLay float64, perc float64) float64 {
	return oddLay / (1 - perc)
}
