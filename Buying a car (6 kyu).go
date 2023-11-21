package main

import "math"

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	var priceOld, priceNew float64
	priceOld, priceNew = float64(startPriceOld), float64(startPriceNew)

	month := 1
	available := priceOld - priceNew

	if priceOld >= priceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}

	for {
		dif := (priceNew - priceOld) / 100 * percentLossByMonth
		available += dif + float64(savingperMonth)
		priceOld -= priceOld*percentLossByMonth/100 - 0.5*float64(savingperMonth)
		priceNew -= priceNew*percentLossByMonth/100 - 0.5*float64(savingperMonth)

		if available >= 0 {
			break
		}

		month++
		if month%2 == 0 {
			percentLossByMonth += 0.5
		}
	}

	return [2]int{
		month,
		int(math.Round(available)),
	}
}
