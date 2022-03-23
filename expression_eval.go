package main

import "fmt"

func EvalExpression(resolver CellResolver, e Expression) (Expression, error) {
	switch concreteExpression := e.(type) {
	case ExpString:
		return concreteExpression, nil
	case ExpList:
		result := make(ExpList, 0)
		for _, elem := range concreteExpression {
			evaluatedExp, err := EvalExpression(resolver, elem)
			if err != nil {
				return nil, fmt.Errorf("err:%w. error evaluating expression:%v", err, concreteExpression)
			}
			result = append(result, evaluatedExp)
		}

		return result, nil
	case ExpFunction:
		evaluatedArgs, err := EvalExpression(resolver, concreteExpression.Args)
		if err != nil {
			return nil, err
		}

		return concreteExpression.Fn(evaluatedArgs)
	case ExpReference:
		cellExpr, err := resolver.GetCell(CellLocation(concreteExpression))
		if err != nil {
			return nil, fmt.Errorf("err:%w. cell:%v", err, concreteExpression)
		}

		return cellExpr, nil
	}

	return nil, ErrUnimplemented
}
