package removeduplicates

var tests = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		expected: []int{0, 1, 2, 3, 4},
	},
	{
		nums:     []int{},
		expected: []int{},
	},
	{
		nums:     []int{0},
		expected: []int{0},
	},
	{
		nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		expected: []int{0},
	},
}
