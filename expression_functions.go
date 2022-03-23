package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Sum(e Expression) (Expression, error) {
	var sum float64
	argList, ok := e.(ExpList)
	if !ok {
		return nil, fmt.Errorf("err:%w. Sum expects inputs to be ExpList, got %v", ErrInvalidArgument, e)
	}
	for _, elem := range argList {
		stringArg, ok := elem.(ExpString)
		if !ok {
			return nil, fmt.Errorf("err:%w. Sum expects elements of ExpList to be string, got %v", ErrInvalidArgument, e)
		}

		v, err := strconv.ParseFloat(string(stringArg), 64)
		if err != nil {
			return nil, fmt.Errorf("err:%w. Sum expects all inputs to be number, got %v", ErrInvalidArgument, elem)
		}

		sum += v
	}

	return ExpString(strconv.FormatFloat(sum, 'f', -1, 64)), nil

}

func Spread(e Expression) (Expression, error) {
	_, ok := e.(ExpList)
	if !ok {
		return nil, fmt.Errorf("err %w. Spread function expects ExpList type, got %t", ErrInvalidArgument, e)
	}

	return e, nil
}

func Split(e Expression) (Expression, error) {
	argList, ok := e.(ExpList)
	if !ok {
		return nil, fmt.Errorf("err %w. Split function expects 2 ExpList type, got %t", ErrInvalidArgument, e)
	}

	if len(argList) != 2 {
		return nil, fmt.Errorf("err %w. Split function expects 2 args, got %d", ErrInvalidArgument, len(argList))
	}

	inputString, ok := argList[0].(ExpString)
	if !ok {
		return nil, fmt.Errorf("err %w. Split function expects 1st arg to be a string, got %t", ErrInvalidArgument, argList[0])
	}

	separator, ok := argList[1].(ExpString)
	if !ok {
		return nil, fmt.Errorf("err %w. Split function expects 2nd arg to be a string, got %t", ErrInvalidArgument, argList[0])
	}

	resultString := strings.Split(string(inputString), string(separator))
	result := make(ExpList, 0)
	for _, elem := range resultString {
		result = append(result, ExpString(strings.TrimSpace(elem)))
	}

	return result, nil
}
