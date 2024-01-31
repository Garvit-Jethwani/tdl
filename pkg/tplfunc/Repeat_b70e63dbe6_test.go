// ********RoostGPT********
/*
Test generated by RoostGPT for test golang-sample-programs using AI Type Open AI and AI Model gpt-4

1. Scenario: Testing the Repeat function with a non-empty string and a positive integer.
Input: ("abc", 3)
Expected Output: "abcabcabc"

2. Scenario: Testing the Repeat function with an empty string and a positive integer.
Input: ("", 5)
Expected Output: ""

3. Scenario: Testing the Repeat function with a non-empty string and zero.
Input: ("abc", 0)
Expected Output: ""

4. Scenario: Testing the Repeat function with a non-empty string and a negative integer.
Input: ("abc", -3)
Expected Output: ""

5. Scenario: Testing the Repeat function with a string that contains special characters and a positive integer.
Input: ("@#%", 5)
Expected Output: "@#%@#%@#%@#%@#%"

6. Scenario: Testing the Repeat function with a string that contains numbers and a positive integer.
Input: ("123", 4)
Expected Output: "123123123123"

7. Scenario: Testing the Repeat function with a string that contains spaces and a positive integer.
Input: (" ab ", 2)
Expected Output: " ab  ab "

8. Scenario: Testing the Repeat function with a string that contains unicode characters and a positive integer.
Input: ("こんにちは", 2)
Expected Output: "こんにちはこんにちは"

9. Scenario: Testing the Repeat function with a very large integer.
Input: ("abc", 1000000)
Expected Output: "abc" repeated 1000000 times

10. Scenario: Testing the Repeat function with a string that contains a newline character and a positive integer.
Input: ("abc\n", 2)
Expected Output: "abc\nabc\n"
*/

// ********RoostGPT********
package tplfunc

import (
	"strings"
	"testing"
	"text/template"
)

func TestRepeat_b70e63dbe6(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		inputStr string
		inputInt int
		expected string
	}{
		{"Non-empty string, positive integer", "abc", 3, strings.Repeat("abc", 3)},
		{"Empty string, positive integer", "", 5, strings.Repeat("", 5)},
		{"Non-empty string, zero", "abc", 0, strings.Repeat("abc", 0)},
		{"Non-empty string, negative integer", "abc", -3, strings.Repeat("abc", -3)},
		{"String with special characters, positive integer", "@#%", 5, strings.Repeat("@#%", 5)},
		{"String with numbers, positive integer", "123", 4, strings.Repeat("123", 4)},
		{"String with spaces, positive integer", " ab ", 2, strings.Repeat(" ab ", 2)},
		{"String with unicode characters, positive integer", "こんにちは", 2, strings.Repeat("こんにちは", 2)},
		{"Very large integer", "abc", 1000000, strings.Repeat("abc", 1000000)},
		{"String with newline character, positive integer", "abc\n", 2, strings.Repeat("abc\n", 2)},
	}

	// Define the function to be tested
	repeatFunc := Repeat()

	funcMap := template.FuncMap{}
	repeatFunc(funcMap)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function and capture the output
			output := funcMap["repeat"].(func(string, int) string)(tc.inputStr, tc.inputInt)

			// Compare the output with the expected result
			if output != tc.expected {
				t.Errorf("Failed on %s: Expected %s but got %s", tc.name, tc.expected, output)
			} else {
				t.Logf("Success on %s", tc.name)
			}
		})
	}
}
