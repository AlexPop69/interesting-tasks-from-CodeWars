package main

import (
	"sort"
	"strings"
)

func FirstNonRepeating(str string) string {
	if str == "" {
		return ""
	}

	newStr := strings.ToLower(str)

	countOfSymbols := make(map[string]int, 0)
	for _, symbol := range newStr {
		countOfSymbols[string(symbol)]++
	}

	idxOfNonRepeatingSymb := make([]int, 0)
	isExistNonRepeatingSymbols := false
	for symbol, countOfRepeat := range countOfSymbols {
		if countOfRepeat == 1 {
			idxOfNonRepeatingSymb = append(idxOfNonRepeatingSymb, strings.Index(newStr, symbol))
			isExistNonRepeatingSymbols = true
		}
	}
	if !isExistNonRepeatingSymbols {
		return ""
	}

	sort.Ints(idxOfNonRepeatingSymb)
	return string(str[idxOfNonRepeatingSymb[0]])
}

/*the best but not my solution
func FirstNonRepeating(str string) string {
    for _, c := range str {
        if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
	          return string(c)
	      }
    }
    return ""
} */
