package removefromlinkedlist

var tests = []struct {
	input    []int
	n        int
	expected []int
}{
	{
		input:    []int{1, 2, 3, 4, 5},
		n:        2,
		expected: []int{1, 2, 3, 5},
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		n:        1,
		expected: []int{1, 2, 3, 4},
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		n:        5,
		expected: []int{2, 3, 4, 5},
	},
}
