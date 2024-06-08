package rm1

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

func GroupAnagrams(strs []string) [][]string {
	// anagram mapped to indices of it's actual words
	return [][]string{}
}
