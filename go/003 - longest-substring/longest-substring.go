package longestsubstring

// package main

func lengthOfLongestSubstring(input string) int {
	inR := []rune(input)
	substring := []rune{}
	for i, r := range inR {
		currentSubstring := []rune{r}
		if i+1 == len(inR) {
			if len(currentSubstring) > len(substring) {
				substring = currentSubstring
			}
			continue
		}
		seenRunes := map[rune]bool{}
		seenRunes[r] = true
		for j := i + 1; j < len(inR); j++ {
			if seenRunes[inR[j]] {
				break
			}
			currentSubstring = append(currentSubstring, inR[j])
			seenRunes[inR[j]] = true
		}
		if len(currentSubstring) > len(substring) {
			substring = currentSubstring
		}
	}
	// fmt.Println(string(substring))
	return len(substring)
}

// package removeduplicates

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	uniqIdx := 0
// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 {
// 			continue
// 		}
// 		if nums[i] != nums[uniqIdx] {
// 			uniqIdx++
// 			nums[uniqIdx] = nums[i]
// 		}
// 	}
// 	return uniqIdx + 1
// }
