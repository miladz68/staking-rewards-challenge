package main

type Expression interface {
	Name() string
	IsEqual(Expression) bool
}

type ExpString string

func (e ExpString) Name() string {
	return "String Expression"
}

func (e ExpString) IsEqual(e1 Expression) bool {
	return e == e1
}

type ExpFunction struct {
	Fn   func(e Expression) (Expression, error)
	Args Expression
}

func (e ExpFunction) Name() string {
	return "Function Expression"
}

func (e ExpFunction) IsEqual(e1 Expression) bool {
	return false
}

type ExpReference CellLocation

func (e ExpReference) Name() string {
	return "Reference Expression"
}

func (e ExpReference) IsEqual(e1 Expression) bool {
	return e == e1
}

type CellLocation struct {
	Column string
	Row    int64
}

type ExpList []Expression

func (e ExpList) Name() string {
	return "List Expression"
}

func (e ExpList) IsEqual(e1 Expression) bool {
	list, ok := e1.(ExpList)
	if !ok {
		return false
	}
	if len(e) != len(list) {
		return false
	}

	for i := range list {
		if e[i] != list[i] {
			return false
		}
	}

	return true
}
