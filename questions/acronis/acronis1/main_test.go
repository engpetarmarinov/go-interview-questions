package main

import "testing"

func TestFizzBuzzNumber(t *testing.T) {
	cases := []struct {
		Name string
		In   int
		Exp  string
	}{
		{
			Name: "5",
			In:   5,
			Exp:  "Buzz",
		},
		{
			Name: "3",
			In:   3,
			Exp:  "Fizz",
		},
		{
			Name: "15",
			In:   15,
			Exp:  "FizzBuzz",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := FizzBuzzNumber(c.In)
			if res != c.Exp {
				t.Errorf("got %s, want %s", res, c.Exp)
			}
		})
	}
}
