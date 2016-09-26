package glob

import "testing"

func TestIsFullMatch(t *testing.T) {
	values := []string{
		"",
		"a",
		"",
		"",
		"", // 5
		"",
		"abbcc",
		"abbcc",
		"abbcc",
		"abbcc", // 10
		"axxxbdxxxc",
		"axxxbdxxxcd",
		"axxxbdxxxd",
		"aaa",
		"azzzzbdbc", // 15
		"azzzzbdbx",
		"abc",
		"zaaz",
		"zaaz",
		"az", // 20
		"za"}
	patterns := []string{
		"",
		"",
		"a",
		"*",
		"***", // 5
		"bc",
		"abbcc",
		"bc",
		"ab?c?",
		"a***z", // 10
		"a*b?*c",
		"a*b?*c",
		"a*b?*c",
		"aa",
		"a***bc", // 15
		"a***bc",
		"**",
		"*a*",
		"*a*",
		"*a*", // 20
		"*a*"}
	results := []bool{
		true,
		false,
		false,
		true,
		true, // 5
		false,
		true,
		false,
		true,
		false, // 10
		true,
		false,
		false,
		false,
		true, // 15
		false,
		true,
		true,
		true,
		true, // 20
		true}

	for i := range patterns {
		if IsFullMatch(values[i], patterns[i]) != results[i] {
			t.Fatalf("Value: %v, Pattern: %v, Expected: %v, Got: %v, Index: %v", values[i], patterns[i], results[i], !results[i], i)
		}
	}
}
