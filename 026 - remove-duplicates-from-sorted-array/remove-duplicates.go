package removeduplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqIdx := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if nums[i] != nums[uniqIdx] {
			uniqIdx++
			nums[uniqIdx] = nums[i]
		}
	}
	return uniqIdx + 1
}
