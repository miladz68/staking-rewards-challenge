package main

type CellResolver interface {
	GetCell(CellLocation) (Expression, error)
}

type StubCellResolver struct {
	state map[CellLocation]Expression
}

func NewStubCellResolver() StubCellResolver {
	return StubCellResolver{
		state: make(map[CellLocation]Expression),
	}
}

func (c StubCellResolver) Add(k CellLocation, v Expression) StubCellResolver {
	c.state[k] = v
	return c
}

func (c StubCellResolver) GetCell(k CellLocation) (Expression, error) {
	value, ok := c.state[k]
	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}
