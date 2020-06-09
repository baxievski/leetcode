package longestsubstring

func lengthOfLongestSubstring(input string) int {
	inR := []rune(input)

	substring := []rune{}
	startIndexes := map[rune]int{}
	i := 0

	for j, r := range inR {
		if val, ok := startIndexes[r]; ok {
			i = max(val, i)
		}

		startIndexes[r] = j + 1

		currentSubstring := inR[i : j+1]
		if len(currentSubstring) > len(substring) {
			substring = currentSubstring
		}
	}

	return len(substring)
}

func max(vars ...int) int {
	m := vars[0]

	for _, i := range vars {
		if m < i {
			m = i
		}
	}

	return m
}
