package utils

func PercPLOpenBack(oddBack float64, oddLay float64) float64 {
	return oddBack/oddLay - 1
}

func PercPLOpenLay(oddLay float64, oddBack float64) float64 {
	return 1 - oddLay/oddBack
}
