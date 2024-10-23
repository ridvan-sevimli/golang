package linkedlist

type LinkedList[T any] struct {
	root *node[T]
}

type node[T any] struct {
	value *T
	next  *node[T]
}

func (l *LinkedList[T]) Init() *LinkedList[T] {
	rootNode := node[T]{value: nil, next: nil}
	return &LinkedList[T]{root: &rootNode}
}

func New[T any]() *LinkedList[T] {
	return new(LinkedList[T]).Init()
}

// Methods

// Add(E element) - Adds an element at the end of the list
func (l *LinkedList[T]) Add(value T) {
	newNode := node[T]{value: &value, next: nil}
	node := l.root
	hasNext := true
	for hasNext {
		if node.next != nil {
			node = node.next
		} else {
			hasNext = false
		}
	}
	node.next = &newNode
}

// AddAt (int index, E value)
func (l *LinkedList[T]) AddAt(index int, value T) {
	if hasNext(l.root) {
		count := 0
		nodez := l.root.next
		for index >= count && hasNext(nodez) {
			prevNode := nodez
			currentNode := nodez.next
			if index == count {
				newNode := node[T]{value: &value, next: currentNode}
				prevNode.next = &newNode
			}
			count++
		}
	} else {
		l.Add(value)
	}
}

// Get(int index)
func (l LinkedList[T]) Get(index int) T {
	var zero T
	if hasNext(l.root) {
		counter := 0
		node := l.root.next
		for index > 0 && hasNext(node) {
			if counter == index {
				return *node.value
			}
			node = node.next
			counter++
		}
		return *node.value
	}
	return zero
}

// Set(index int, E element) - Replaces the element that the specified poistion with the specifeid element
func (l *LinkedList[T]) Set(index int, value T) bool {
	if hasNext(l.root) {
		counter := 0
		node := l.root.next
		for index > 0 && hasNext(node) {
			if counter == index {
				node.value = &value
				return true
			}
			node = node.next
			counter++
		}
	}
	return false
}

// RemoveFirst()
func (l *LinkedList[T]) RemoveFirst() {
	if hasNext(l.root) {
		l.root.next = l.root.next.next
	}
}

// RemoveLast()
func (l *LinkedList[T]) RemoveLast() {
	if hasNext(l.root) {
		node := l.root.next
		for hasNext(node) {
			node = node.next
		}
		node.next = nil
	}
}

// GetFirst()
func (l *LinkedList[T]) GetFirst() T {
	var zero T
	if hasNext(l.root) {
		return *l.root.next.value
	}
	return zero
}

// GetLast()
func (l *LinkedList[T]) GetLast() T {
	var zero T
	if hasNext(l.root) {
		node := l.root.next
		for hasNext(node) {
			node = node.next
		}
		return *node.value
	}
	return zero
}

// Remove(int index) - Removes the element at the specified index

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

func hasNext[T any](node *node[T]) bool {
	return node.next != nil
}
