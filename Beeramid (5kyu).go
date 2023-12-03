package main

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 {
		return 0
	}
	line := 1.0
	cans := 1.0
	referralBonus := float64(bonus)

	for {
		if (cans * price) > referralBonus {
			line--
		}
		referralBonus -= cans * price
		if referralBonus <= 0 {
			break
		}
		line++
		cans = (line * line)
	}
	return int(line)
}
