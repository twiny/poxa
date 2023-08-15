package poxa

import "container/ring"

type (
	Spinner[T any] interface {
		Next() T
	}

	spinner[T any] struct {
		elems *ring.Ring
	}
)

func NewSpinner[T any](elems ...T) Spinner[T] {
	if len(elems) == 0 {
		return nil
	}

	s := &spinner[T]{
		elems: ring.New(len(elems)),
	}

	for _, elem := range elems {
		s.elems.Value = elem
		s.elems = s.elems.Next()
	}

	return s
}

func (s *spinner[T]) Next() T {
	if s.elems == nil {
		var zero T
		return zero
	}

	elem := s.elems.Value
	s.elems = s.elems.Next()

	return elem.(T)
}
