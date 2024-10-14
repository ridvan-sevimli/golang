package list

type List[T any] struct {
	values []T
}

func (l *List[T]) Init() *List[T] {
	return &List[T]{values: make([]T, 0)}
}

func New[T any]() *List[T] {
	return new(List[T]).Init()
}

// Methods

// add(E element) - Adds an element at the end of the lsit
// addAtIndex (int index, E element)
// get(int index)
// set(int index, E element) - Replaces the element tat the specified poistion with the specifeid element
// removeFirst()
// removeLast()
// getFirst()
// getLast()
// remove(int index) - Removes the element at the specified index
// remove(Object o) - Removes the first occurrence of the specified element
// peek() - Retrieves, but does not remove, the head of the list, or teturns null if the list is empty

// Traversal Method

// traverse()

// Search

// search(value)

// Update Method

// updateAtIndex(int index,E element)

//Utility Methods

// contains(Object o) - Returns true if the list contains the specified element.
// clear() - Removes all elements from the list.
// size() - Returns the number of elements in the list.
// indexOf(Object o)- Returns the index of the first occurence of the speciefied element
// isEmpty()
// reverse()
// lastIndexOf(Object o) - Returns the index of the last occurrence of the specified element
