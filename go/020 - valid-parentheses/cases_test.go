package validparentheses

var tests = []struct {
	input    string
	expected bool
}{
	{
		input:    "()",
		expected: true,
	},
	{
		input:    "[]",
		expected: true,
	},
	{
		input:    "{}",
		expected: true,
	},
	{
		input:    "()[]{}",
		expected: true,
	},
	{
		input:    "(]",
		expected: false,
	},
	{
		input:    "([)]",
		expected: false,
	},
	{
		input:    "{[()]}",
		expected: true,
	},
}
