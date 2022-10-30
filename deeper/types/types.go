package types

type Interface interface {
	Method()
}
type String struct {
	Value string
}
type Integer int

func (s *String) Method() {}

func (i Integer) Method() {}
