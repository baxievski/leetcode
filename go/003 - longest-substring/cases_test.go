package longestsubstring

var tests = []struct {
	input    string
	expected int
}{
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
}
