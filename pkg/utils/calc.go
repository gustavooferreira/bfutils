package utils

func CalcPercOpenBack(oddBack float64, oddLay float64) float64 {
	return oddBack/oddLay - 1
}

func CalcPercOpenLay(oddLay float64, oddBack float64) float64 {
	return 1 - oddLay/oddBack
}
