package longestcommonprefix

var tests = []struct {
	description string
	inputs      []string
	expected    string
}{
	{
		description: "three ascii strings",
		inputs:      []string{"flower", "flow", "flight"},
		expected:    "fl",
	},
	{
		description: "three strings, cyrilic ж at start of first",
		inputs:      []string{"жflower", "flow", "flight"},
		expected:    "",
	},
	{
		description: "three strings, cyrilic ж at start of second",
		inputs:      []string{"flower", "жflow", "flight"},
		expected:    "",
	},
	{
		description: "three strings, cyrilic ж at start of third",
		inputs:      []string{"flower", "flow", "жflight"},
		expected:    "",
	},
	{
		description: "three strings, cyrilic ж at end of third",
		inputs:      []string{"flower", "flow", "flightж"},
		expected:    "fl",
	},
	{
		description: "three strings, cyrilic ж at start of all",
		inputs:      []string{"жflower", "жflow", "жflight"},
		expected:    "жfl",
	},
	{
		description: "three strings, first empty",
		inputs:      []string{"", "жflow", "жflight"},
		expected:    "",
	},
	{
		description: "three strings, all empty",
		inputs:      []string{"", "", ""},
		expected:    "",
	},
	{
		description: "one empty string",
		inputs:      []string{""},
		expected:    "",
	},
	{
		description: "no strings ",
		inputs:      []string{},
		expected:    "",
	},
	{
		description: "three cyrilic strings",
		inputs:      []string{"Бојан", "БојанАџиевски", "Бојанаџиевски"},
		expected:    "Бојан",
	},
}
