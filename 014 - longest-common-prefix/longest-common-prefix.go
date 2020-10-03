package longestcommonprefix

func longestCommonPrefix(inputs []string) string {
	result := []rune{}

	minLen := 0
	inputsR := map[int][]rune{}
	for i, s := range inputs {
		inputsR[i] = []rune(s)
		if i == 0 {
			minLen = len(inputsR[0])
			continue
		}
		if len(inputsR[i]) < minLen {
			minLen = len(inputsR[i])
		}
	}

	for i := 0; i < minLen; i++ {
		curRune := inputsR[0][i]
		for j := 0; j < len(inputs); j++ {
			if inputsR[j][i] != curRune {
				return string(result)
			}
		}
		result = append(result, curRune)
	}

	return string(result)
}
