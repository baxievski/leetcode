package twosum

func twoSum(nums []int, target int) []int {
	numsToIndexes := map[int]int{}
	for i, num := range nums {
		remainder := target - num
		if j, ok := numsToIndexes[remainder]; ok {
			return []int{j, i}
		}
		numsToIndexes[num] = i
	}
	return []int{}
}
