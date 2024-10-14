package linkedlist

import "testing"

func TestAdd(t *testing.T) {
	linkedList := New[int]()

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.addAt(1, 17)
	got := linkedList.get(3)
	want := 3
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
