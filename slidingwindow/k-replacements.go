package slidingwindow

import (
	"fmt"
)

func getMostFrequentRune(letterFrequencies map[rune]int) int { // Iterating over the map
	maxVal := 0
	for _, val := range letterFrequencies {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
func kReplacements() {
	s := "ABAB"
	k := 2
	letterFrequencies := make(map[rune]int)
	l, r, maxCount := 0, 0, 0

	for r < len(s) {
		// With the current letterFrequencies, get the most frequent word and (r - l) which is the current
		// length of the substring. (r-l) - (mostFrequent) <= k, you can continue adding to the substring
		letterFrequencies[rune(s[r])] += 1
		mostFrequentCount := getMostFrequentRune(letterFrequencies)
		windowLength := (r - l) + 1
		// fmt.Printf("Most frequent (%d)\nl: %d, r: %d\nmaxCount: %d\n\n", mostFrequentCount, l, r, maxCount)

		// Is this a valid substring?
		if (windowLength - mostFrequentCount) <= k {
			// This length can be considered for the maxCount
			maxCount = max(windowLength, maxCount)
		} else {
			letterFrequencies[rune(s[l])] -= 1
			l++
		}

		r++
	}

	fmt.Println(maxCount)
}
