package main

import (
	"errors"
	"testing"
)

func TestParseExpression(t *testing.T) {
	testCases := []struct {
		input              string
		ExpectedExpression Expression
		ExpectedError      error
	}{
		{
			input: "=sum(1,2,3)",
			ExpectedExpression: ExpFunction{
				Fn:   Sum,
				Args: ExpList{ExpString("1"), ExpString("2"), ExpString("3")},
			},
		},
		{
			input: `=sum(spread(split(D2, ",")))`,
			ExpectedExpression: ExpFunction{
				Fn: Sum,
				Args: ExpFunction{
					Fn: Spread,
					Args: ExpFunction{
						Fn: Split,
						Args: ExpList{
							ExpReference{Column: "D", Row: 2},
							ExpString(","),
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		got, err := ParseExpression(test.input)
		if !errors.Is(test.ExpectedError, err) {
			t.Log("got unexpected error", err)
			t.FailNow()
		}

		if got != test.ExpectedExpression {
			t.Log("wrong expression parsed", err)
			t.FailNow()
		}
	}
}
