package main

import "testing"

func TestMinimize(t *testing.T) {
	cin1 := `
	/*
	 * Function to chop a string in half.
	*/
	public static string chop(string input) {
		if (input == null || input.isEmpty()) {
			return input;
		}
		if (input.length() % 2 == 1) {
			return "cannot chop an odd-length string in half";
		}
		return input.substring(input.length() / 2);
	}`

	cout1 := `
	/*
	 * Function to chop a string in half.
	*/
	public static $4 $2($4 input) {
		if ($12 == null || $12.isEmpty()) {
			return $12;
		}
		$13 ($12.length() % 2 == 1) {
			$18 "cannot $2 an odd-$22 $4 $5 $6";
		}
		$18 $12.substring($12.$22() / 2);
	}`
	var inputTests = []struct {
		name string
		in   string
		out  string
	}{
		{"s0", "", ""},
		{"s1", "abc", "abc"},
		{"s2", "you you you", "you $0 $0"},
		{"s3", "you say yes, I say no you say stop and I say go go go", "you say yes, I $1 no $0 $1 stop and $3 $1 go $12 $12"},
		{"c1", cin1, cout1},
	}

	for _, tt := range inputTests {
		t.Run(tt.name, func(t *testing.T) {
			m := minimize(tt.in)
			if m != tt.out {
				t.Errorf("Got %q, want %q", m, tt.out)
			}
		})
	}
}
