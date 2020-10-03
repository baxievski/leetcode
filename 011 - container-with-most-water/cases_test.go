package containerwithmostwater

var tests = []struct {
	inputs   []int
	expected int
}{
	{
		inputs:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		inputs:   []int{2, 3, 10, 5, 7, 8, 9},
		expected: 36,
	},
}
