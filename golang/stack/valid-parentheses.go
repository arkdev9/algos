package stack

import "fmt"

type StackNode struct {
	val  rune
	next *StackNode
}

func (s *StackNode) push(val rune) *StackNode {
	// Create a new node
	newNode := &StackNode{val: val, next: s}
	return newNode
}

func (s *StackNode) read() {
	// Print out stack from top to bottom
	if s == nil {
		fmt.Println("Empty stack")
	}
	for s != nil {
		fmt.Println(s.val)
		s = s.next
	}
}

func (s *StackNode) pop() *StackNode {
	return s.next
}

func IsValid(s string) bool {
	var head *StackNode
	runes := []rune(s)
	bracketMap := map[rune]rune{'}': '{', ']': '[', ')': '('}

	fmt.Println(bracketMap)
	fmt.Println(runes)

	// Iterate through the source string as runes
	for i := 0; i < len(runes); i++ {
		head.read()
		_, exists := bracketMap[rune(runes[i])]
		currRune := runes[i]
		if !exists {
			// this is an opening bracket, push it into the stack
			head = head.push(runes[i])

		} else {
			// it's a closing bracket, pop from the stack
			// if stack is empty, return invalid, because not possible to close
			if head == nil {
				return false
			}

			peekedRune := head.val
			// peekedRune has to be a matching opening bracket
			if peekedRune == bracketMap[currRune] {
				head = head.pop()
			} else {
				return false
			}
		}
	}
	head.read()
	// After running through all runes, stack should be empty
	return head == nil
}
