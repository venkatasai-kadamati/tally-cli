package algo

import (
	"fmt"
	"strings"
)

// CountWords Func: implement algorithm for counting words in a file
// idea: notice space to increase counter
func CountWords(data string) int {
	// considerable cases
	// 1. manually its really an extended effort to get the additional spaces considered at start/end. Sol: trim the spaces around.
	// 2. look further on how empty text files work in different encoding.
	fmt.Println("data received:", data)

	// if data (FILE) is empty
	if len(data) == 0 {
		return 0
	}

	maxStreak := 1
	checker := " "

	// trim the spaces
	data = strings.TrimSpace(data)
	for _, val := range data {
		if string(val) == checker {
			fmt.Println("noticed checker")
			maxStreak++
		}
	}

	// handling the missing word at the last.
	// issue: normally since the end
	//if len(data) > 1 {
	//	maxStreak++
	//}

	fmt.Println("count :", maxStreak)
	return maxStreak
}

// notes: we can use empty char (runes : utf-8)
// introducing new cases to test behaviour in code is called as regression.
// to control regressions, we use automated testing
// we have 3 famous automated tests: Unit, E2E, Integrated
// best split of above is 70:20:10 (unit, integrated, e2e)
