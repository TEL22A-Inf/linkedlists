package dllist_int

import "fmt"

// A doubly linked list.
// Contains a pointer to the anchor element.
type DLList struct {
	anchor *DLLNode
}

// NewDLList creates a new empty doubly linked list.
func NewDLList() *DLList {
	l := &DLList{}
	l.anchor = NewDLLNode()
	return l
}

// IsEmpty returns true if the list is empty.
func (l *DLList) IsEmpty() bool {
	return l.anchor.IsEmpty()
}

// PushBack adds a new element at the end of the list.
func (l *DLList) PushBack(value int) {
	l.anchor.Prev.Insert(value)
}

// PushFront adds a new element at the beginning of the list.
func (l *DLList) PushFront(value int) {
	l.anchor.Insert(value)
}

// PopBack removes the last element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopBack() {
	l.anchor.Prev.Remove()
}

// PopFront removes the first element from the list.
// If the list is empty, nothing happens.
func (l *DLList) PopFront() {
	l.anchor.Next.Remove()
}

// String returns a string representation of the list.
func (l *DLList) String() string {
	valuestrings := []string{}
	for e := l.anchor.Next; e != l.anchor; e = e.Next {
		valuestrings = append(valuestrings, e.String())
	}
	return fmt.Sprintf("%v", valuestrings)
}
