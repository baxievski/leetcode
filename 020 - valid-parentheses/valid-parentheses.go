package validparentheses

func isValid(input string) bool {
	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	inR := []rune(input)
	parentheses := []rune{}

	for _, curParen := range inR {
		if openParen, ok := matchingBrackets[curParen]; ok {
			if len(parentheses) == 0 {
				return false
			}

			if parentheses[len(parentheses)-1] != openParen {
				return false
			}

			parentheses = parentheses[:len(parentheses)-1]
			continue
		}

		parentheses = append(parentheses, curParen)
	}

	return len(parentheses) == 0
}
