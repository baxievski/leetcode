package longestsubstring

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "aaaaaa",
		expected: 1,
	},
	{
		input:    "abcabcbb",
		expected: 3,
	},
	{
		input:    "abcabcdabcdea",
		expected: 5,
	},
	{
		input:    "",
		expected: 0,
	},
	{
		input:    "аабабвабвг",
		expected: 4,
	},
	{
		input:    " ",
		expected: 1,
	},
	{
		input:    "abc",
		expected: 3,
	},
}
