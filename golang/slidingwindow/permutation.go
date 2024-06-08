package slidingwindow

import "fmt"

func createStringSignature(s string) string {
	runeRun := [26]int{}
	for i := 0; i < len(s); i++ {
		runeRun[int(rune(s[i]))-97] += 1
	}

	outString := ""

	for _, val := range runeRun {
		outString += "#" + fmt.Sprint(val)
	}
	return outString
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Sig := createStringSignature(s1)
	fmt.Println(s1Sig)
	// window length can be fixed to size of s1
	l, r := 0, len(s1)

	for r <= len(s2) {
		// The window of r-l should have the same string signature for it to be a permutation
		if s1Sig == createStringSignature(s2[l:r]) {
			return true
		}

		l++
		r++
	}

	return false
}
