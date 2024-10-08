package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	stack := New[int]() // Create new Stack
	stack.Push(1)       // add value to stack
	stack.Push(2)
	stack.Push(3)
	got := *stack                             // derefrence stack
	want := Stack[int]{items: []int{1, 2, 3}} // value to check against

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestPeek(t *testing.T) {
	t.Run("Stack has elements", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		got, err := stack.Peek()
		want := 3

		if err != nil {
			t.Fatal("Did not expect an error")
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Stack is empty", func(t *testing.T) {

		stackEmpty := New[int]()
		got, err := stackEmpty.Peek()
		var want int

		if err != stackEmptyError {
			t.Fatalf("Expected error")
		}
		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}

	})

}

func TestPop(t *testing.T) {
	t.Run("Stack has element returns value", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		got, err := stack.Pop()
		want := 3

		if err != nil {
			t.Fatal("Did not expect an error")
		}

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("Check if poped element was removed from the Stack", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Pop()
		got := stack.items
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

}

func TestRemoveAll(t *testing.T) {
	stack := New[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.RemoveAll()

	got := stack.items
	want := make([]int, 0)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestIsEmpty(t *testing.T) {
	t.Run("Stack has elements", func(t *testing.T) {
		stack := New[int]()

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		got := stack.IsEmpty()
		want := false

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("Stack is empty", func(t *testing.T) {
		stack := New[int]()

		got := stack.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("Check Pop every item", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Pop()
		stack.Pop()
		stack.Pop()
		got := stack.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
