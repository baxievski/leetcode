package twosum

var tests = []struct {
	description string
	nums        []int
	target      int
	expected    []int
}{
	{
		description: "First two elements sum to target",
		nums:        []int{2, 7, 11, 15},
		target:      9,
		expected:    []int{0, 1},
	},
	{
		description: "Somewhere in the middle",
		nums:        []int{2, 7, 11, 15, 8, 12, 13, 90, 22, 32, 44},
		target:      19,
		expected:    []int{2, 4},
	},
	{
		description: "First and last in slice of len 2",
		nums:        []int{3, 3},
		target:      6,
		expected:    []int{0, 1},
	},
	{
		description: "First and last in slice of len 3",
		nums:        []int{3, 1, 3},
		target:      6,
		expected:    []int{0, 2},
	},
	{
		description: "First and last, one is 0",
		nums:        []int{6, 1, 3, 0},
		target:      6,
		expected:    []int{0, 3},
	},
	{
		description: "None",
		nums:        []int{2, 1, 3, 0},
		target:      60,
		expected:    []int{},
	},
	{
		description: "None from empty slice",
		nums:        []int{},
		target:      60,
		expected:    []int{},
	},
}
