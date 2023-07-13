package stack

func IsValid(s string) bool {
	parenStack := make([]rune)

	for i := 0; i < len(s); i++ {
		currRune := s[i]
		// If it is an opening bracket, push onto stack
		if currRune == "(" || currRune == "{" || currRune == "[" {

		}

		parenStack.append(parenStack, currRune)
	}

	return true
}
