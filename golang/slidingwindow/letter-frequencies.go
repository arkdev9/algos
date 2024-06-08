package slidingwindow

import (
	"fmt"
)

func max(a, b int) int {
	// Returns b if it is greater than a
	if b > a {
		return b
	}
	return a
}

func lengthOfLongestSubstring(s string) int {
	letterMap := make(map[rune]int)
	maxWindowLength, l := 0, 0

	for r := 0; r < len(s); r++ {
		currRune := rune(s[r])
		letterMap[currRune] += 1
		windowLength := r - l + 1
		fmt.Printf("currRune: %c\nwindowLength: %d\nl: %d, r: %d\n\n", currRune, windowLength, l, r)
		// Get current count of this rune
		count := letterMap[currRune]
		if count > 1 {
			for rune(s[l]) != currRune {
				letterMap[rune(s[l])] -= 1
				l += 1
			}
			letterMap[rune(s[l])] -= 1
			l += 1
		} else {
			maxWindowLength = max(windowLength, maxWindowLength)
		}
	}

	return maxWindowLength
}
