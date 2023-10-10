package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		excepted []string
	}{
		{input: "Hello World",
			excepted: []string{
				"hello",
				"world",
			},
		},
		{input: "HELLO World",
			excepted: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := clearInput(cs.input)
		if len(actual) != len(cs.excepted) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(actual),
				len(cs.excepted),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.excepted[i]
			if actualWord != expectedWord {
				t.Errorf("%v doesn't equal %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}
}
