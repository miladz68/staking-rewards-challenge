package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestEvalExpression(t *testing.T) {
	testCases := []struct {
		resolver CellResolver
		input    Expression
		result   Expression
		err      error
	}{
		{
			input: ExpFunction{
				Fn:   Sum,
				Args: ExpList{ExpString("1"), ExpString("2"), ExpString("3")},
			},
			result: ExpString("6"),
		},
		{
			input: ExpFunction{
				Fn:   Sum,
				Args: ExpList{ExpString("1.1"), ExpString("2.1"), ExpString("3")},
			},
			result: ExpString("6.2"),
		},
		{
			input:    ExpReference{Column: "D", Row: 2},
			result:   ExpString("1.1, 3.1, 5.1"),
			resolver: NewStubCellResolver().Add(CellLocation{"D", 2}, ExpString("1.1, 3.1, 5.1")),
		},
		{
			input: ExpFunction{
				Fn: Split,
				Args: ExpList{
					ExpReference{Column: "D", Row: 2},
					ExpString(","),
				},
			},
			result: ExpList{ExpString("1.1"), ExpString("3.1"), ExpString("5.1")},
			resolver: NewStubCellResolver().
				Add(CellLocation{"D", 2}, ExpString("1.1, 3.1, 5.1")),
		},
		{
			input: ExpFunction{
				Fn: Sum,
				Args: ExpFunction{
					Fn: Split,
					Args: ExpList{
						ExpReference{Column: "D", Row: 2},
						ExpString(","),
					},
				},
			},
			result: ExpString("9.3"),
			resolver: NewStubCellResolver().
				Add(CellLocation{"D", 2}, ExpString("1.1, 3.1, 5.1")),
		},
		{
			input: ExpFunction{
				Fn:   Spread,
				Args: ExpList{ExpString("1.1"), ExpString("3.1"), ExpString("5.1")},
			},
			result: ExpList{ExpString("1.1"), ExpString("3.1"), ExpString("5.1")},
		},
		{
			input: ExpFunction{
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
			result: ExpString("9.3"),
			resolver: NewStubCellResolver().
				Add(CellLocation{"D", 2}, ExpString("1.1, 3.1, 5.1")),
		},
	}

	for _, test := range testCases {
		got, err := EvalExpression(test.resolver, test.input)
		if !errors.Is(test.err, err) {
			t.Log("got unexpected error", err)
			t.FailNow()
		}

		if !got.IsEqual(test.result) {
			fmt.Println(got, test.result)
			t.Log("wrong expression evaluated", err)
			t.FailNow()
		}
	}
}
