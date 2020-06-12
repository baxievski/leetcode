package validparentheses

import "fmt"

func isValid(input string) bool {
	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '}',
	}
	inR := []rune(input)
	parentheses := []rune{}
	for i, curParen := range inR {
		fmt.Println(i, string(curParen), string(parentheses))
		if openParen, ok := matchingBrackets[curParen]; ok {
			if parentheses[len(parentheses)-1] != openParen {
				return false
			}
			parentheses = parentheses[:len(parentheses)-1]
		}
		parentheses = append(parentheses, curParen)
	}

	if len(parentheses) == 0 {
		return false
	}
	return true
}
