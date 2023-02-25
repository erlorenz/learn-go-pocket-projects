package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {

	type testCase struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]testCase{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour le monde",
		},
		"Akkadian": {
			lang:             "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Unsupported": {
			lang:             "",
			expectedGreeting: `unsupported language: ""`,
		},
	}

	// range over all the tests
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
