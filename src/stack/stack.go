package stack

import "errors"

var stackEmptyError error = errors.New("Stack is empty")

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Init() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func New[T any]() *Stack[T] {
	return new(Stack[T]).Init()
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, stackEmptyError
	}
	defer func() { s.items = s.items[:len(s.items)-1] }()

	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, stackEmptyError
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) RemoveAll() {
	newStack := new(Stack[T]).Init()
	*s = *newStack
}
